package nextmail

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"nextmail.sh/internal/protocol"
)

const (
	optionAPIKey         = "apiKey"
	optionIdempotencyKey = "idempotencyKey"
	optionInboxToken     = "inboxToken"
	optionTimeout        = "timeoutMs"
	optionQueryParam     = "queryParam"
	optionLimit          = "limit"
	optionCursor         = "cursor"
	withAPIKey           = "WithAPIKey"
	withIdempotencyKey   = "WithIdempotencyKey"
	withInboxToken       = "WithInboxToken"
	withTimeout          = "WithTimeout"
	withQueryParam       = "WithQueryParam"
	rawRequestAccessor   = "Request"
)

var rawRequestOptions = []string{withAPIKey, withIdempotencyKey, withInboxToken, withQueryParam}

type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}

type ClientOption interface {
	applyClient(config *clientConfig)
}

type RequestOption interface {
	applyRequest(options *callOptions)
}

type Option interface {
	ClientOption
	RequestOption
}

type header struct {
	name  string
	value string
}

type clientConfig struct {
	apiKey        secret
	hasAPIKey     bool
	accessToken   secret
	hasToken      bool
	tokenFunc     func(ctx context.Context) (string, error)
	inboxToken    secret
	baseURL       string
	hasBaseURL    bool
	httpClient    HTTPClient
	maxRetries    int
	hasMaxRetries bool
	timeout       time.Duration
	hasTimeout    bool
	userAgent     string
	headers       []header
	err           error
}

func (c *clientConfig) fail(err error) {
	if c.err == nil {
		c.err = err
	}
}

type clientOption func(config *clientConfig)

func (o clientOption) applyClient(config *clientConfig) {
	o(config)
}

type callOption struct {
	function string
	name     string
	value    any
}

func (o callOption) applyRequest(options *callOptions) {
	options.set(o.function, o.name, o.value)
}

type secretOption struct {
	function string
	name     string
	value    secret
}

func (o secretOption) applyRequest(options *callOptions) {
	options.set(o.function, o.name, o.value)
}

func (o secretOption) applyClient(config *clientConfig) {
	switch o.name {
	case optionAPIKey:
		config.apiKey = o.value
		config.hasAPIKey = true
	case optionInboxToken:
		config.inboxToken = o.value
	}
}

func (o secretOption) String() string {
	return "nextmail." + o.function + "(" + redacted + ")"
}

func (o secretOption) GoString() string {
	return o.String()
}

func (o secretOption) Format(state fmt.State, _ rune) {
	_, _ = fmt.Fprint(state, o.String())
}

type timeoutOption time.Duration

func (o timeoutOption) applyRequest(options *callOptions) {
	if o < 0 {
		options.fail(argumentError(messageTimeoutShape))

		return
	}

	options.set(withTimeout, optionTimeout, time.Duration(o))
}

func (o timeoutOption) applyClient(config *clientConfig) {
	if o < 0 {
		config.fail(argumentError(messageTimeoutShape))

		return
	}

	config.timeout = time.Duration(o)
	config.hasTimeout = true
}

func WithAPIKey(apiKey string) Option {
	return secretOption{function: withAPIKey, name: optionAPIKey, value: newSecret(apiKey)}
}

func WithInboxToken(inboxToken string) Option {
	return secretOption{function: withInboxToken, name: optionInboxToken, value: newSecret(inboxToken)}
}

func WithTimeout(timeout time.Duration) Option {
	return timeoutOption(timeout)
}

func WithAccessToken(accessToken string) ClientOption {
	token := newSecret(accessToken)

	return clientOption(func(config *clientConfig) {
		config.accessToken = token
		config.tokenFunc = nil
		config.hasToken = true
	})
}

func WithAccessTokenFunc(provider func(ctx context.Context) (string, error)) ClientOption {
	return clientOption(func(config *clientConfig) {
		if provider == nil {
			config.fail(argumentError(messageAccessTokenFunc))

			return
		}

		config.accessToken = secret{}
		config.tokenFunc = provider
		config.hasToken = true
	})
}

func WithBaseURL(baseURL string) ClientOption {
	return clientOption(func(config *clientConfig) {
		config.baseURL = baseURL
		config.hasBaseURL = true
	})
}

func WithHTTPClient(client HTTPClient) ClientOption {
	return clientOption(func(config *clientConfig) {
		if client == nil {
			config.fail(argumentError(messageHTTPClientRequired))

			return
		}

		config.httpClient = client
	})
}

func WithMaxRetries(maxRetries int) ClientOption {
	return clientOption(func(config *clientConfig) {
		config.maxRetries = max(0, maxRetries)
		config.hasMaxRetries = true
	})
}

func WithUserAgent(userAgent string) ClientOption {
	return clientOption(func(config *clientConfig) {
		config.userAgent = userAgent
	})
}

func WithHeader(name, value string) ClientOption {
	return clientOption(func(config *clientConfig) {
		config.headers = append(config.headers, header{name: name, value: value})
	})
}

func WithIdempotencyKey(idempotencyKey string) RequestOption {
	return callOption{function: withIdempotencyKey, name: optionIdempotencyKey, value: idempotencyKey}
}

func WithQueryParam(name, value string) RequestOption {
	return callOption{function: withQueryParam, name: optionQueryParam, value: pair{key: name, value: value}}
}

type callOptions struct {
	accessor string
	allowed  []string
	values   map[string]any
	extra    values
	err      error
}

func newCallOptions(accessor string, opts []RequestOption) *callOptions {
	options := &callOptions{accessor: accessor, values: map[string]any{}}

	if accessor == rawRequestAccessor {
		options.allowed = rawRequestOptions
	} else if allowed, known := methodOptions[accessor]; known {
		options.allowed = allowed
	} else {
		options.fail(argumentError(messageUnknownMethod, accessor))
	}

	for _, option := range opts {
		if option == nil {
			options.fail(argumentError(messageNilOption, accessor))

			continue
		}

		option.applyRequest(options)
	}

	return options
}

func (o *callOptions) fail(err error) {
	if o.err == nil {
		o.err = err
	}
}

func (o *callOptions) set(function, name string, value any) {
	if function != withTimeout && !slices.Contains(o.allowed, function) {
		accepted := messageNoOptions

		if len(o.allowed) > 0 {
			accepted = strings.Join(o.allowed, ", ")
		}

		o.fail(argumentError(messageOptionNotAccepted, function, o.accessor, accepted))

		return
	}

	if entry, ok := value.(pair); ok {
		o.extra = append(o.extra, entry)

		return
	}

	o.values[name] = value
}

func (o *callOptions) value(name string) any {
	switch typed := o.values[name].(type) {
	case time.Time:
		return instant(typed)
	case ImageCrop:
		return typed.text()
	case []string:
		return strings.Join(typed, protocol.ListSeparator)
	case []int:
		texts := make([]string, len(typed))

		for index, number := range typed {
			texts[index] = strconv.Itoa(number)
		}

		return strings.Join(texts, protocol.ListSeparator)
	default:
		return typed
	}
}

func (o *callOptions) text(name string) string {
	text, _ := o.value(name).(string)

	return text
}

func (o *callOptions) has(name string) bool {
	_, present := o.values[name]

	return present
}

func (o *callOptions) secret(name string) secret {
	value, _ := o.values[name].(secret)

	return value
}

func (o *callOptions) timeout() (time.Duration, bool) {
	value, present := o.values[optionTimeout].(time.Duration)

	return value, present
}
