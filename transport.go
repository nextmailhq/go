package nextmail

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	mathrand "math/rand/v2"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"nextmail.sh/internal/protocol"
	"nextmail.sh/internal/query"
)

const (
	userAgentPrefix    = "nextmail-go/"
	bearerPrefix       = "Bearer "
	errorMessageLength = 300
	httpWhitespace     = " \t\r\n"
	byteOrderMark      = "\xEF\xBB\xBF"
	statusOKFloor      = 200
	statusOKCeiling    = 300
)

var errNoResponse = errors.New("the HTTP client returned no response")

var (
	headerNamePattern      = regexp.MustCompile("^[!#$%&'*+\\-.^_`|~0-9A-Za-z]+$")
	headerForbiddenPattern = regexp.MustCompile(`[\x00-\x08\x0A-\x1F\x7F]`)
	decimalPattern         = regexp.MustCompile(`^[+-]?(?:\d+\.?\d*(?:[eE][+-]?\d+)?|\.\d+(?:[eE][+-]?\d+)?)$`)
)

type pair struct {
	key   string
	value any
}

type values []pair

func (q values) with(key string, value any) values {
	for index := range q {
		if q[index].key == key {
			updated := slices.Clone(q)
			updated[index].value = value

			return updated
		}
	}

	return append(slices.Clone(q), pair{key: key, value: value})
}

func (q values) merge(other values) values {
	merged := q

	for _, entry := range other {
		merged = merged.with(entry.key, entry.value)
	}

	return merged
}

type cursorStyle struct {
	param string
	field string
}

var (
	cursorPages    = cursorStyle{param: protocol.CursorStylesCursorParam, field: protocol.CursorStylesCursorField}
	pageTokenPages = cursorStyle{param: protocol.CursorStylesPageTokenParam, field: protocol.CursorStylesPageTokenField}
)

type call struct {
	method      string
	path        string
	params      []any
	query       values
	body        any
	raw         io.Reader
	sendsRaw    bool
	contentType string
	anonymous   bool
	idempotent  bool
	repeatable  bool
	accept      string
	upload      bool
	turn        bool
}

type reply struct {
	status int
	header http.Header
	body   []byte
}

func (r *reply) ok() bool {
	return r.status >= statusOKFloor && r.status < statusOKCeiling
}

type transport struct {
	baseURL     string
	secure      bool
	mode        string
	apiKey      secret
	accessToken secret
	inboxToken  secret
	tokenFunc   func(ctx context.Context) (string, error)
	client      HTTPClient
	maxRetries  int
	timeout     time.Duration
	userAgent   string
	headers     func() []header
	err         error
	sleep       func(ctx context.Context, wait time.Duration) error
	now         func() time.Time
	jitter      func() float64
}

func newTransport(config *clientConfig, credentialRequired bool) *transport {
	t := &transport{
		baseURL:     protocol.DefaultsBaseURL,
		mode:        protocol.APIKeyModesLive,
		apiKey:      config.apiKey,
		accessToken: config.accessToken,
		inboxToken:  config.inboxToken,
		tokenFunc:   config.tokenFunc,
		client:      config.httpClient,
		maxRetries:  protocol.DefaultsMaxRetries,
		timeout:     protocol.DefaultsTimeoutMs * time.Millisecond,
		userAgent:   userAgentPrefix + Version,
		headers:     func() []header { return nil },
		sleep:       pause,
		now:         time.Now,
		jitter:      mathrand.Float64,
	}

	if t.client == nil {
		t.client = &http.Client{CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
	}

	if config.hasMaxRetries {
		t.maxRetries = config.maxRetries
	}

	if config.hasTimeout {
		t.timeout = config.timeout
	}

	if strings.HasPrefix(t.apiKey.value(), protocol.APIKeyPrefixesTest) {
		t.mode = protocol.APIKeyModesTest
	}

	t.err = t.configure(config, credentialRequired)

	return t
}

func (t *transport) configure(config *clientConfig, credentialRequired bool) error {
	if config.err != nil {
		return config.err
	}

	if err := assertCredential(t.apiKey, t.accessToken, t.tokenFunc != nil, credentialRequired); err != nil {
		return err
	}

	if config.hasBaseURL || config.baseURL != "" {
		baseURL, err := validBaseURL(config.baseURL)

		if err != nil {
			return err
		}

		t.baseURL = baseURL
	}

	t.secure = isSecureOrigin(t.baseURL)

	if config.userAgent != "" {
		userAgent, err := headerValue(config.userAgent)

		if err != nil {
			return err
		}

		t.userAgent = userAgent
	}

	custom := make([]header, 0, len(config.headers))

	for _, entry := range config.headers {
		value, err := headerValue(entry.value)

		if err != nil || !headerNamePattern.MatchString(entry.name) {
			return argumentError(messageHeaderShape)
		}

		custom = append(custom, header{name: entry.name, value: value})
	}

	t.headers = func() []header { return custom }

	return nil
}

func headerValue(value string) (string, error) {
	trimmed := strings.Trim(value, httpWhitespace)

	if headerForbiddenPattern.MatchString(trimmed) || !utf8.ValidString(trimmed) {
		return "", argumentError(messageHeaderShape)
	}

	return trimmed, nil
}

func pause(ctx context.Context, wait time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if wait <= 0 {
		return nil
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func (t *transport) backoff(attempt int) time.Duration {
	ceiling := math.Min(protocol.RetryMaxBackoffMs, protocol.RetryBaseBackoffMs*math.Pow(2, float64(attempt)))

	return time.Duration(math.Round(ceiling*(0.5+t.jitter()*0.5))) * time.Millisecond
}

func idempotencyKey() string {
	var raw [16]byte

	_, _ = rand.Read(raw[:])
	raw[6] = raw[6]&0x0f | 0x40
	raw[8] = raw[8]&0x3f | 0x80

	text := hex.EncodeToString(raw[:])

	return protocol.IdempotencyKeyPrefix + text[0:8] + "-" + text[8:12] + "-" + text[12:16] + "-" + text[16:20] + "-" + text[20:32]
}

func retryAfterSeconds(value string, now time.Time) (float64, bool) {
	text := strings.TrimSpace(value)

	if text == "" {
		return 0, false
	}

	if decimalPattern.MatchString(text) {
		seconds, err := strconv.ParseFloat(text, 64)

		if err == nil && !math.IsInf(seconds, 0) && !math.IsNaN(seconds) && seconds >= 0 {
			return seconds, true
		}

		return 0, false
	}

	at, err := http.ParseTime(text)

	if err != nil {
		at, err = time.Parse(time.RFC3339Nano, text)
	}

	if err != nil {
		return 0, false
	}

	return math.Max(0, math.Ceil(at.Sub(now).Seconds())), true
}

func (t *transport) execute(ctx context.Context, options *callOptions, request call) (*reply, error) {
	if t.err != nil {
		return nil, t.err
	}

	if options.err != nil {
		return nil, options.err
	}

	if ctx == nil {
		ctx = context.Background()
	}

	method := strings.ToUpper(request.method)

	if method == "" {
		method = http.MethodGet
	}

	if !headerNamePattern.MatchString(method) {
		return nil, argumentError(messageMethodShape)
	}

	path, err := fillPath(request.path, request.params)

	if err != nil {
		return nil, err
	}

	target, err := buildURL(t.baseURL, path, request.query.merge(options.extra))

	if err != nil {
		return nil, err
	}

	headers, err := t.headersFor(ctx, options, request)

	if err != nil {
		return nil, err
	}

	body, err := bodyBytes(request)

	if err != nil {
		return nil, err
	}

	repeatable := request.repeatable || slices.Contains(protocol.RetryMethods, method)
	timeout := t.timeoutFor(options, request)

	for attempt := 0; ; attempt++ {
		answer, failure := t.attempt(ctx, method, target, headers, body, timeout)

		if failure != nil {
			if failure.kind != kindNetwork || ctx.Err() != nil || !repeatable || attempt >= t.maxRetries {
				return nil, failure
			}

			if err := t.sleep(ctx, t.backoff(attempt)); err != nil {
				return nil, networkError(err, false, messageAborted)
			}

			continue
		}

		if answer.ok() {
			return answer, nil
		}

		seconds, given := retryAfterSeconds(answer.header.Get(protocol.HeaderKeysRetryAfter), t.now())
		honoured := answer.status != protocol.RetryRateLimitStatus

		if given {
			honoured = seconds*1000 <= protocol.RetryMaxHonouredRetryAfterMs
		}

		if repeatable && honoured && attempt < t.maxRetries && slices.Contains(protocol.RetryStatuses, answer.status) {
			wait := t.backoff(attempt)

			if given {
				wait = time.Duration(seconds * float64(time.Second))
			}

			if err := t.sleep(ctx, wait); err != nil {
				return nil, networkError(err, false, messageAborted)
			}

			continue
		}

		return nil, apiError(answer, seconds, given)
	}
}

func (t *transport) timeoutFor(options *callOptions, request call) time.Duration {
	if chosen, given := options.timeout(); given {
		return chosen
	}

	if request.upload && t.timeout > 0 {
		return max(t.timeout, protocol.DefaultsUploadTimeoutMs*time.Millisecond)
	}

	if request.turn && t.timeout > 0 {
		return max(t.timeout, protocol.ChatTurnTimeoutMs*time.Millisecond)
	}

	return t.timeout
}

func bodyBytes(request call) ([]byte, error) {
	if request.sendsRaw {
		if request.raw == nil {
			return nil, argumentError(messageRawBodyRequired)
		}

		data, err := io.ReadAll(request.raw)

		if err != nil {
			return nil, argumentFailure(err, messageRawBodyUnreadable, err.Error())
		}

		if data == nil {
			data = []byte{}
		}

		return data, nil
	}

	if request.body == nil {
		return nil, nil
	}

	return encodeBody(request.body)
}

func (t *transport) headersFor(ctx context.Context, options *callOptions, request call) ([]header, error) {
	headers := slices.Clone(t.headers())
	accept := request.accept

	if accept == "" {
		accept = protocol.ContentTypesJSON
	}

	headers = append(headers, header{name: protocol.HeaderKeysAccept, value: accept}, header{name: protocol.HeaderKeysUserAgent, value: t.userAgent})

	if request.sendsRaw {
		contentType := request.contentType

		if contentType == "" {
			contentType = protocol.ContentTypesOctetStream
		}

		headers = append(headers, header{name: protocol.HeaderKeysContentType, value: contentType})
	} else if request.body != nil {
		headers = append(headers, header{name: protocol.HeaderKeysContentType, value: protocol.ContentTypesJSON})
	}

	if !request.anonymous {
		credential, err := t.credentialFor(ctx, options)

		if err != nil {
			return nil, err
		}

		if credential != "" {
			headers = append(headers, header{name: protocol.HeaderKeysAuthorization, value: bearerPrefix + credential})
		}
	}

	if request.idempotent {
		key := options.text(optionIdempotencyKey)

		if key == "" {
			key = idempotencyKey()
		}

		headers = append(headers, header{name: protocol.HeaderKeysIdempotencyKey, value: key})
	}

	for index := range headers {
		value, err := headerValue(headers[index].value)

		if err != nil {
			return nil, err
		}

		headers[index].value = value
	}

	return headers, nil
}

func (t *transport) credentialFor(ctx context.Context, options *callOptions) (string, error) {
	if options.has(optionAPIKey) {
		if err := assertAPIKey(options.secret(optionAPIKey).value(), messageCallAPIKeySubject); err != nil {
			return "", err
		}
	}

	direct, chosen := t.directCredential(options)
	withToken := t.accessToken.present() || t.tokenFunc != nil

	if (direct != "" || (!chosen && withToken)) && !t.secure {
		return "", argumentError(messageCleartext, t.baseURL)
	}

	if chosen {
		return direct, nil
	}

	if t.tokenFunc == nil {
		return t.accessToken.value(), nil
	}

	token, err := t.tokenFunc(ctx)

	if err != nil {
		return "", argumentFailure(err, messageAccessTokenFailed, err.Error())
	}

	if !IsAccessToken(token) {
		return "", argumentError("%s %s", messageAccessTokenProvided, messageAccessTokenShape)
	}

	return token, nil
}

func (t *transport) directCredential(options *callOptions) (string, bool) {
	switch {
	case options.has(optionInboxToken):
		return options.secret(optionInboxToken).value(), true
	case options.has(optionAPIKey):
		return options.secret(optionAPIKey).value(), true
	case t.inboxToken.present():
		return t.inboxToken.value(), true
	case t.apiKey.present():
		return t.apiKey.value(), true
	}

	return "", false
}

func (t *transport) attempt(ctx context.Context, method, target string, headers []header, body []byte, timeout time.Duration) (*reply, *Error) {
	deadline, release := context.WithCancel(ctx)

	if timeout > 0 {
		deadline, release = context.WithTimeout(ctx, timeout)
	}

	defer release()

	var payload io.Reader

	if body != nil {
		payload = bytes.NewReader(body)
	}

	request, err := http.NewRequestWithContext(deadline, method, target, payload)

	if err != nil {
		return nil, argumentError("%s %s", quoted(target[len(t.baseURL):]), messagePathShape)
	}

	for _, entry := range headers {
		request.Header.Set(entry.name, entry.value)
	}

	failure := func(cause error, fallback string) *Error {
		if ctx.Err() != nil {
			return networkError(errors.Join(ctx.Err(), cause), false, messageAborted)
		}

		if deadline.Err() != nil {
			return networkError(errors.Join(deadline.Err(), cause), true, messageTimedOut, timeout)
		}

		return networkError(cause, false, "%s: %s", fallback, describe(cause))
	}

	response, err := t.client.Do(request)

	if err != nil {
		return nil, failure(err, fmt.Sprintf(messageUnreachable, t.baseURL))
	}

	if response == nil {
		return nil, failure(errNoResponse, fmt.Sprintf(messageUnreachable, t.baseURL))
	}

	answer := &reply{status: response.StatusCode, header: response.Header}

	if response.Body == nil {
		return answer, nil
	}

	defer func() { _ = response.Body.Close() }()

	answer.body, err = io.ReadAll(response.Body)

	if err != nil {
		if answer.ok() || ctx.Err() != nil || deadline.Err() != nil {
			return nil, failure(err, messageBodyUnreadable)
		}

		answer.body = nil
	}

	return answer, nil
}

func describe(err error) string {
	var wrapped *url.Error

	if errors.As(err, &wrapped) && wrapped.Err != nil {
		return wrapped.Err.Error()
	}

	return err.Error()
}

func decodeJSON(text []byte) (any, bool) {
	var decoded any

	if err := json.Unmarshal(text, &decoded); err != nil {
		return nil, false
	}

	return decoded, true
}

func textOf(body []byte) string {
	return strings.TrimPrefix(strings.ToValidUTF8(string(body), replacementCharacter), byteOrderMark)
}

func apiError(answer *reply, seconds float64, given bool) *Error {
	text := textOf(answer.body)
	body, _ := decodeJSON([]byte(text))
	payload := objectOf(objectOf(body)["error"])
	failure := &Error{
		Message:   payload.String("message"),
		Status:    answer.status,
		Type:      payload.String("type"),
		Code:      payload.String("code"),
		Param:     payload.String("param"),
		DocURL:    payload.String("docUrl"),
		RequestID: payload.String("requestId"),
		Body:      body,
		kind:      kindAPI,
	}

	if failure.Message == "" {
		failure.Message = fallbackMessage(answer.status, text)
	}

	if !slices.Contains(errorTypes, failure.Type) {
		failure.Type = typeForStatus(answer.status)
	}

	if failure.Code == "" {
		failure.Code = protocol.ErrorCodesUnrecognisedResponse
	}

	if failure.RequestID == "" {
		failure.RequestID = answer.header.Get(protocol.HeaderKeysRequestID)
	}

	if given {
		failure.RetryAfter = time.Duration(seconds * float64(time.Second))
	}

	if fields, listed := payload["fields"].([]any); listed {
		failure.Fields = []FieldProblem{}

		for _, entry := range fields {
			field := objectOf(entry)
			key, hasKey := field["key"].(string)
			problem, hasProblem := field["error"].(string)

			if hasKey && hasProblem {
				failure.Fields = append(failure.Fields, FieldProblem{Key: key, Error: problem})
			}
		}
	}

	return failure
}

var errorTypes = []string{
	protocol.ErrorTypesInvalidRequestError,
	protocol.ErrorTypesAuthenticationError,
	protocol.ErrorTypesPermissionError,
	protocol.ErrorTypesNotFoundError,
	protocol.ErrorTypesConflictError,
	protocol.ErrorTypesValidationError,
	protocol.ErrorTypesRateLimitError,
	protocol.ErrorTypesAPIError,
}

func typeForStatus(status int) string {
	if status >= http.StatusInternalServerError {
		return protocol.ErrorTypesAPIError
	}

	if mapped, known := protocol.ErrorTypeByStatus[status]; known && slices.Contains(errorTypes, mapped) {
		return mapped
	}

	return protocol.ErrorTypesInvalidRequestError
}

func fallbackMessage(status int, text string) string {
	if text == "" {
		return fmt.Sprintf(messageUnrecognisedBody, status)
	}

	runes := []rune(text)

	if len(runes) > errorMessageLength {
		return string(runes[:errorMessageLength])
	}

	return text
}

func unrecognised(answer *reply, message string, body any) *Error {
	return &Error{
		Message:   message,
		Status:    answer.status,
		Type:      protocol.ErrorTypesAPIError,
		Code:      protocol.ErrorCodesUnrecognisedResponse,
		RequestID: answer.header.Get(protocol.HeaderKeysRequestID),
		Body:      body,
		kind:      kindAPI,
	}
}

func (t *transport) json(ctx context.Context, options *callOptions, request call) (any, *reply, error) {
	answer, err := t.execute(ctx, options, request)

	if err != nil {
		return nil, nil, err
	}

	text := textOf(answer.body)

	if text == "" {
		return nil, answer, nil
	}

	decoded, ok := decodeJSON([]byte(text))

	if !ok {
		return nil, nil, unrecognised(answer, fmt.Sprintf(messageNotJSON, answer.status), nil)
	}

	return decoded, answer, nil
}

func (t *transport) envelope(ctx context.Context, options *callOptions, request call) (Object, error) {
	decoded, answer, err := t.json(ctx, options, request)

	if err != nil {
		return nil, err
	}

	if decoded == nil {
		return Object{}, nil
	}

	object := objectOf(decoded)

	if object == nil {
		return nil, unrecognised(answer, messageNotJSONObject, decoded)
	}

	return object, nil
}

func (t *transport) object(ctx context.Context, options *callOptions, request call) (Object, error) {
	return t.envelope(ctx, options, request)
}

func (t *transport) objects(ctx context.Context, options *callOptions, request call) ([]Object, error) {
	envelope, err := t.envelope(ctx, options, request)

	if err != nil {
		return nil, err
	}

	return objectsOf(envelope[fieldData]), nil
}

func (t *transport) text(ctx context.Context, options *callOptions, request call) (string, error) {
	answer, err := t.execute(ctx, options, request)

	if err != nil {
		return "", err
	}

	return textOf(answer.body), nil
}

func (t *transport) bytes(ctx context.Context, options *callOptions, request call) ([]byte, error) {
	answer, err := t.execute(ctx, options, request)

	if err != nil {
		return nil, err
	}

	if answer.body == nil {
		return []byte{}, nil
	}

	return answer.body, nil
}

func (t *transport) pageAt(ctx context.Context, options *callOptions, request call, style cursorStyle, cursor string) (*Page, error) {
	request.query = request.query.with(query.Limit, options.value(optionLimit)).with(style.param, cursor)

	envelope, err := t.envelope(ctx, options, request)

	if err != nil {
		return nil, err
	}

	next := cursorOf(envelope[style.field])

	return &Page{Items: objectsOf(envelope[fieldData]), HasMore: hasMoreOf(envelope, next), NextCursor: next}, nil
}

func (t *transport) page(ctx context.Context, options *callOptions, request call, style cursorStyle) (*Page, error) {
	return t.pageAt(ctx, options, request, style, options.text(optionCursor))
}

func (t *transport) iterate(ctx context.Context, options *callOptions, request call, style cursorStyle) *Iterator {
	return newIterator(ctx, options.text(optionCursor), func(ctx context.Context, cursor string) ([]Object, bool, string, error) {
		page, err := t.pageAt(ctx, options, request, style, cursor)

		if err != nil {
			return nil, false, "", err
		}

		return page.Items, page.HasMore, page.NextCursor, nil
	})
}

func (t *transport) all(ctx context.Context, options *callOptions, request call, style cursorStyle) ([]Object, error) {
	return t.iterate(ctx, options, request, style).collect()
}
