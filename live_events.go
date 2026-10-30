package nextmail

import (
	"context"
	"errors"
	"io"
	"iter"
	"slices"
	"sync"
	"time"

	"nextmail.sh/constants"
	"nextmail.sh/internal/paths"
	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

const (
	LiveEventReset  = protocol.LiveStreamEventsReset
	fieldType       = "type"
	fieldOccurredAt = "occurredAt"
	fieldReason     = "reason"
)

type LiveEvent struct {
	ID         string
	Type       string
	OccurredAt string
	Data       Object
}

type LiveEventStream struct {
	ctx        context.Context
	cancel     context.CancelFunc
	t          *transport
	options    *callOptions
	request    call
	timing     liveTiming
	resumeID   string
	retry      time.Duration
	failures   int
	delay      time.Duration
	started    time.Time
	mutex      sync.Mutex
	connection *streamConnection
	parser     *eventParser
	event      LiveEvent
	err        error
	done       bool
}

func (s *LiveEventsService) Stream(ctx context.Context, opts ...RequestOption) *LiveEventStream {
	o := newCallOptions("liveEvents.stream", opts)

	return newLiveEventStream(ctx, s.t, o, call{
		path:   paths.LiveEvents,
		query:  values{{query.Types, o.value("types")}, {query.Address, o.value("address")}},
		accept: protocol.ContentTypesEventStream,
	}, o.text("lastEventId"))
}

func newLiveEventStream(ctx context.Context, t *transport, options *callOptions, request call, resumeID string) *LiveEventStream {
	if ctx == nil {
		ctx = context.Background()
	}

	scope, cancel := context.WithCancel(ctx)
	timing := t.live

	return &LiveEventStream{
		ctx:      scope,
		cancel:   cancel,
		t:        t,
		options:  options,
		request:  request,
		timing:   timing,
		resumeID: resumeID,
		retry:    min(timing.defaultRetry, timing.maxBackoff),
	}
}

func (s *LiveEventStream) Next() bool {
	if s == nil || s.done {
		return false
	}

	for {
		if s.parser == nil {
			if !s.connect() {
				return false
			}

			continue
		}

		record, err := s.parser.next()

		if err != nil {
			s.disconnect(!errors.Is(err, io.EOF))

			if s.ctx.Err() != nil {
				return s.finish(nil)
			}

			continue
		}

		if record.hasRetry {
			s.retry = min(record.retry, s.timing.maxBackoff)
		}

		if record.hasID {
			s.resumeID = record.id
		}

		if !record.hasData {
			continue
		}

		if record.event == protocol.LiveStreamEventsEnd {
			s.disconnect(false)

			if s.ctx.Err() != nil {
				return s.finish(nil)
			}

			continue
		}

		if event, delivered := s.eventOf(record); delivered {
			s.event = event

			return true
		}
	}
}

func (s *LiveEventStream) connect() bool {
	if err := s.t.sleep(s.ctx, s.delay); err != nil || s.ctx.Err() != nil {
		return s.finish(nil)
	}

	s.started = s.t.now()
	connection, err := s.t.stream(s.ctx, s.options, s.attempt(), s.timing.idle)

	if err != nil {
		if s.ctx.Err() != nil {
			return s.finish(nil)
		}

		var failure *Error

		if !errors.As(err, &failure) || !reconnects(failure) {
			return s.finish(err)
		}

		s.delay = s.retryDelay(failure.RetryAfter)
		s.failures++

		return true
	}

	if mediaTypeOf(connection.header) != protocol.ContentTypesEventStream {
		connection.close()

		return s.finish(notAnEventStream(connection))
	}

	s.failures = 0
	s.mutex.Lock()
	s.connection = connection
	s.mutex.Unlock()
	s.parser = newEventParser(connection)

	return true
}

func (s *LiveEventStream) attempt() call {
	request := s.request

	if s.resumeID != "" {
		request.headers = []header{{name: protocol.HeaderKeysLastEventID, value: s.resumeID}}
	}

	return request
}

func reconnects(failure *Error) bool {
	if failure.kind == kindNetwork {
		return true
	}

	return failure.kind == kindAPI && (failure.Retryable() || failure.Status >= 500)
}

func (s *LiveEventStream) retryDelay(retryAfter time.Duration) time.Duration {
	backoff := s.t.backoffWithin(s.failures, milliseconds(s.timing.baseBackoff), milliseconds(s.timing.maxBackoff))

	if retryAfter <= 0 {
		return backoff
	}

	return max(backoff, min(retryAfter, s.timing.maxRetryAfter))
}

func milliseconds(duration time.Duration) float64 {
	return float64(duration) / float64(time.Millisecond)
}

func (s *LiveEventStream) disconnect(failed bool) {
	s.mutex.Lock()
	connection := s.connection
	s.connection = nil
	s.mutex.Unlock()

	if connection != nil {
		connection.close()
	}

	s.parser = nil

	if failed {
		s.delay = s.retryDelay(0)
		s.failures++

		return
	}

	s.delay = max(0, s.started.Add(s.retry).Sub(s.t.now()))
	s.failures = 0
}

func (s *LiveEventStream) eventOf(record eventRecord) (LiveEvent, bool) {
	decoded, _ := decodeJSON([]byte(record.data))
	payload := objectOf(decoded)

	if record.event == protocol.LiveStreamEventsReset {
		reason, given := payload[fieldReason].(string)

		if !given {
			reason = constants.LiveResetReasonsUnknown
		}

		return LiveEvent{ID: s.resumeID, Type: LiveEventReset, OccurredAt: instant(s.t.now()), Data: Object{fieldReason: reason}}, true
	}

	if payload == nil || !slices.Contains(constants.LiveEventTypesValues(), record.event) {
		return LiveEvent{}, false
	}

	return LiveEvent{ID: payload.String(fieldID), Type: payload.String(fieldType), OccurredAt: payload.String(fieldOccurredAt), Data: payload.Object(fieldData)}, true
}

func (s *LiveEventStream) finish(err error) bool {
	s.done = true
	s.err = err
	s.event = LiveEvent{}
	s.cancel()

	return false
}

func (s *LiveEventStream) Event() LiveEvent {
	if s == nil {
		return LiveEvent{}
	}

	return s.event
}

func (s *LiveEventStream) Err() error {
	if s == nil {
		return nil
	}

	return s.err
}

func (s *LiveEventStream) Close() error {
	if s == nil {
		return nil
	}

	s.cancel()
	s.mutex.Lock()
	connection := s.connection
	s.mutex.Unlock()

	if connection != nil {
		connection.close()
	}

	return nil
}

func (s *LiveEventStream) All() iter.Seq2[LiveEvent, error] {
	return func(yield func(LiveEvent, error) bool) {
		defer func() { _ = s.Close() }()

		for s.Next() {
			if !yield(s.Event(), nil) {
				return
			}
		}

		if err := s.Err(); err != nil {
			yield(LiveEvent{}, err)
		}
	}
}
