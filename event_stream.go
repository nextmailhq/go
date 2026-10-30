package nextmail

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"nextmail.sh/internal/protocol"
)

var (
	errStreamIdle        = errors.New("the event stream sent nothing for too long")
	errStreamOpenTimeout = errors.New("the event stream did not open in time")
)

type liveTiming struct {
	baseBackoff   time.Duration
	maxBackoff    time.Duration
	defaultRetry  time.Duration
	maxRetryAfter time.Duration
	idle          time.Duration
}

func defaultLiveTiming() liveTiming {
	return liveTiming{
		baseBackoff:   protocol.LiveStreamTimingBaseBackoffMs * time.Millisecond,
		maxBackoff:    protocol.LiveStreamTimingMaxBackoffMs * time.Millisecond,
		defaultRetry:  protocol.LiveStreamTimingDefaultRetryMs * time.Millisecond,
		maxRetryAfter: protocol.LiveStreamTimingMaxRetryAfterMs * time.Millisecond,
		idle:          protocol.LiveStreamTimingIdleTimeoutMs * time.Millisecond,
	}
}

type eventRecord struct {
	id       string
	hasID    bool
	event    string
	data     string
	hasData  bool
	retry    time.Duration
	hasRetry bool
}

type eventParser struct {
	source  *bufio.Reader
	line    []byte
	skipLF  bool
	started bool
	block   eventRecord
	data    []string
	stored  int
	touched bool
}

func newEventParser(source io.Reader) *eventParser {
	return &eventParser{source: bufio.NewReader(source)}
}

func (p *eventParser) next() (eventRecord, error) {
	for {
		line, err := p.readLine()

		if err != nil {
			return eventRecord{}, err
		}

		if record, dispatched := p.take(line); dispatched {
			return record, nil
		}
	}
}

func (p *eventParser) readLine() (string, error) {
	p.line = p.line[:0]

	for {
		character, err := p.source.ReadByte()

		if err != nil {
			return "", err
		}

		if p.skipLF {
			p.skipLF = false

			if character == '\n' {
				continue
			}
		}

		if character == '\n' || character == '\r' {
			p.skipLF = character == '\r'

			return p.text(), nil
		}

		p.line = append(p.line, character)

		if p.stored+len(p.line) > protocol.EventStreamSyntaxMaxBlockChars {
			return "", networkError(nil, false, messageEventTooLong, protocol.EventStreamSyntaxMaxBlockChars)
		}
	}
}

func (p *eventParser) text() string {
	line := string(p.line)

	if !p.started {
		p.started = true
		line = strings.TrimPrefix(line, byteOrderMark)
	}

	return strings.ToValidUTF8(line, replacementCharacter)
}

func (p *eventParser) take(line string) (eventRecord, bool) {
	if line == "" {
		record, touched := p.block, p.touched

		if record.event == "" {
			record.event = protocol.EventStreamSyntaxDefaultEvent
		}

		if len(p.data) > 0 {
			record.data = strings.Join(p.data, protocol.EventStreamSyntaxLineFeed)
			record.hasData = true
		}

		p.block = eventRecord{}
		p.data = nil
		p.stored = 0
		p.touched = false

		return record, touched
	}

	if strings.HasPrefix(line, protocol.EventStreamSyntaxColon) {
		return eventRecord{}, false
	}

	name, value, found := strings.Cut(line, protocol.EventStreamSyntaxColon)

	if found {
		value = strings.TrimPrefix(value, protocol.EventStreamSyntaxSpace)
	}

	p.field(name, value)

	return eventRecord{}, false
}

func (p *eventParser) field(name, value string) {
	switch name {
	case protocol.EventStreamFieldsEvent:
		p.block.event = value
	case protocol.EventStreamFieldsData:
		p.data = append(p.data, value)
		p.stored += len(value)
	case protocol.EventStreamFieldsID:
		if strings.Contains(value, protocol.EventStreamSyntaxNull) {
			return
		}

		p.block.id = value
		p.block.hasID = true
	case protocol.EventStreamFieldsRetry:
		retry, valid := retryOf(value)

		if !valid {
			return
		}

		p.block.retry = retry
		p.block.hasRetry = true
	default:
		return
	}

	p.touched = true
}

func retryOf(value string) (time.Duration, bool) {
	if value == "" {
		return 0, false
	}

	for _, character := range value {
		if character < '0' || character > '9' {
			return 0, false
		}
	}

	milliseconds, err := strconv.ParseInt(value, 10, 64)

	if err != nil || milliseconds > math.MaxInt64/int64(time.Millisecond) {
		return time.Duration(math.MaxInt64), true
	}

	return time.Duration(milliseconds) * time.Millisecond, true
}

type streamConnection struct {
	status   int
	header   http.Header
	body     io.ReadCloser
	idle     time.Duration
	watchdog *time.Timer
	cancel   context.CancelCauseFunc
	once     sync.Once
}

func (c *streamConnection) Read(buffer []byte) (int, error) {
	if c.watchdog != nil {
		c.watchdog.Reset(c.idle)
	}

	read, err := c.body.Read(buffer)

	if c.watchdog != nil {
		c.watchdog.Stop()
	}

	return read, err
}

func (c *streamConnection) close() {
	c.once.Do(func() {
		if c.watchdog != nil {
			c.watchdog.Stop()
		}

		c.cancel(nil)
		_ = c.body.Close()
	})
}

func mediaTypeOf(header http.Header) string {
	media, _, _ := strings.Cut(header.Get(protocol.HeaderKeysContentType), protocol.MediaTypeParameterSeparator)

	return strings.ToLower(strings.TrimSpace(media))
}

func notAnEventStream(connection *streamConnection) *Error {
	media := mediaTypeOf(connection.header)

	if media == "" {
		media = messageNoContentType
	}

	return unrecognised(&reply{status: connection.status, header: connection.header}, fmt.Sprintf(messageNotEventStream, connection.status, media), nil)
}

func (t *transport) stream(ctx context.Context, options *callOptions, request call, idle time.Duration) (*streamConnection, error) {
	method, target, headers, err := t.prepare(ctx, options, request)

	if err != nil {
		return nil, err
	}

	timeout := t.timeoutFor(options, request)
	scope, cancel := context.WithCancelCause(ctx)
	outgoing, err := http.NewRequestWithContext(scope, method, target, nil)

	if err != nil {
		cancel(nil)

		return nil, argumentError("%s %s", quoted(target[len(t.baseURL):]), messagePathShape)
	}

	for _, entry := range headers {
		outgoing.Header.Set(entry.name, entry.value)
	}

	var opening *time.Timer

	if timeout > 0 {
		opening = time.AfterFunc(timeout, func() { cancel(errStreamOpenTimeout) })
	}

	disarmed := func() bool {
		return opening == nil || opening.Stop()
	}

	failure := func(cause error, fallback string) *Error {
		defer cancel(nil)

		if ctx.Err() != nil {
			return networkError(errors.Join(ctx.Err(), cause), false, messageAborted)
		}

		if errors.Is(context.Cause(scope), errStreamOpenTimeout) {
			return networkError(errors.Join(context.DeadlineExceeded, cause), true, messageTimedOut, timeout)
		}

		return networkError(cause, false, "%s: %s", fallback, describe(cause))
	}

	response, err := t.client.Do(outgoing)

	if err == nil && response == nil {
		err = errNoResponse
	}

	if err != nil {
		disarmed()

		return nil, failure(err, fmt.Sprintf(messageUnreachable, t.baseURL))
	}

	body := response.Body

	if body == nil {
		body = http.NoBody
	}

	answer := &reply{status: response.StatusCode, header: response.Header}

	if answer.header == nil {
		answer.header = http.Header{}
	}

	if !answer.ok() {
		answer.body, _ = io.ReadAll(body)
		_ = body.Close()
		disarmed()
		cancel(nil)

		seconds, given := retryAfterSeconds(answer.header.Get(protocol.HeaderKeysRetryAfter), t.now())

		return nil, apiError(answer, seconds, given)
	}

	if !disarmed() {
		_ = body.Close()

		return nil, failure(errStreamOpenTimeout, fmt.Sprintf(messageUnreachable, t.baseURL))
	}

	connection := &streamConnection{status: answer.status, header: answer.header, body: body, idle: idle, cancel: cancel}

	if idle > 0 {
		connection.watchdog = time.AfterFunc(idle, func() { cancel(errStreamIdle) })
		connection.watchdog.Stop()
	}

	context.AfterFunc(scope, connection.close)

	return connection, nil
}
