// Code generated from the TypeScript SDK. DO NOT EDIT.

package protocol

const AccessTokenRulesReservedPrefix = "nm_"
const AccessTokenRulesMaxLength = 512

const APIKeyModesLive = "live"
const APIKeyModesTest = "test"

const APIKeyPrefixesLive = "nm_live_"
const APIKeyPrefixesTest = "nm_test_"

const ChatTurnTimeoutMs = 300000

const ContentTypesJSON = "application/json"
const ContentTypesCalendar = "text/calendar"
const ContentTypesVcard = "text/vcard"
const ContentTypesOctetStream = "application/octet-stream"
const ContentTypesEventStream = "text/event-stream"

const CursorStylesCursorParam = "cursor"
const CursorStylesCursorField = "nextCursor"
const CursorStylesPageTokenParam = "pageToken"
const CursorStylesPageTokenField = "nextPageToken"

const DefaultsBaseURL = "https://api.nextmail.sh"
const DefaultsMaxRetries = 2
const DefaultsTimeoutMs = 30000
const DefaultsUploadTimeoutMs = 600000

const EnvVarsAPIKey = "NEXTMAIL_API_KEY"
const EnvVarsAccessToken = "NEXTMAIL_ACCESS_TOKEN"
const EnvVarsBaseURL = "NEXTMAIL_BASE_URL"

const ErrorCodesInsufficientScope = "insufficient_scope"
const ErrorCodesConsoleOnly = "console_only"
const ErrorCodesUnrecognisedResponse = "unrecognised_response"

const ErrorTypesInvalidRequestError = "invalid_request_error"
const ErrorTypesAuthenticationError = "authentication_error"
const ErrorTypesPermissionError = "permission_error"
const ErrorTypesNotFoundError = "not_found_error"
const ErrorTypesConflictError = "conflict_error"
const ErrorTypesValidationError = "validation_error"
const ErrorTypesRateLimitError = "rate_limit_error"
const ErrorTypesAPIError = "api_error"

var ErrorTypeByStatus = map[int]string{
	401: "authentication_error",
	403: "permission_error",
	404: "not_found_error",
	409: "conflict_error",
	422: "validation_error",
	429: "rate_limit_error",
}

const EventStreamFieldsEvent = "event"
const EventStreamFieldsData = "data"
const EventStreamFieldsID = "id"
const EventStreamFieldsRetry = "retry"

const EventStreamSyntaxDefaultEvent = "message"
const EventStreamSyntaxColon = ":"
const EventStreamSyntaxSpace = " "
const EventStreamSyntaxNull = "\u0000"
const EventStreamSyntaxLineFeed = "\n"
const EventStreamSyntaxCarriageReturn = "\r"
const EventStreamSyntaxMaxBlockChars = 1048576

const HeaderKeysAuthorization = "Authorization"
const HeaderKeysAccept = "Accept"
const HeaderKeysContentType = "Content-Type"
const HeaderKeysUserAgent = "User-Agent"
const HeaderKeysIdempotencyKey = "Idempotency-Key"
const HeaderKeysRetryAfter = "retry-after"
const HeaderKeysRequestID = "x-request-id"
const HeaderKeysLastEventID = "Last-Event-ID"

const IdempotencyKeyPrefix = "nm-"

const ImageCropDecimals = 6

const ImageCropSeparator = ","

var LanguageAliases = [][]string{{"zh", "zh-Hans"}, {"zh-cn", "zh-Hans"}, {"zh-sg", "zh-Hans"}, {"cmn", "zh-Hans"}, {"zh-tw", "zh-Hant"}, {"zh-mo", "zh-HK"}, {"iw", "he"}, {"in", "id"}, {"ji", "yi"}, {"mo", "ro"}, {"sh", "sr-Latn"}, {"sr-latn-rs", "sr-Latn"}, {"prs", "fa-AF"}, {"fa-af", "fa-AF"}, {"kmr", "ku"}, {"pt-pt", "pt"}, {"es-mx", "es-419"}, {"es-ar", "es-419"}, {"es-co", "es-419"}, {"es-cl", "es-419"}}

const ListSeparator = ","

const LiveStreamEventsReady = "ready"
const LiveStreamEventsReset = "reset"
const LiveStreamEventsEnd = "end"

const LiveStreamTimingBaseBackoffMs = 1000
const LiveStreamTimingMaxBackoffMs = 30000
const LiveStreamTimingDefaultRetryMs = 3000
const LiveStreamTimingMaxRetryAfterMs = 300000
const LiveStreamTimingIdleTimeoutMs = 60000

var LocalHosts = []string{"localhost", "::1", "[::1]", "127.0.0.1", "0.0.0.0", "::", "[::]"}

var LoopbackHostnames = []string{"localhost", "::1", "[::1]"}

const LoopbackIpv4Address = "127.0.0.1"

const LoopbackIpv6Address = "[::1]"

const MediaTypeParameterSeparator = ";"

var RetryStatuses = []int{408, 429, 500, 502, 503, 504}
var RetryMethods = []string{"GET"}

const RetryBaseBackoffMs = 500
const RetryMaxBackoffMs = 8000
const RetryMaxHonouredRetryAfterMs = 60000
const RetryRateLimitStatus = 429

var SealedEncryptionFormats = []string{"pgp-mime", "pgp-inline", "smime-encrypted"}

const StepUpErrorCodesStepUpRequired = "step_up_required"
const StepUpErrorCodesStepUpNotApplicable = "step_up_not_applicable"
const StepUpErrorCodesStepUpThrottled = "step_up_throttled"
const StepUpErrorCodesStepUpUndeliverable = "step_up_undeliverable"
const StepUpErrorCodesStepUpCodeInvalid = "step_up_code_invalid"
const StepUpErrorCodesStepUpCodeExpired = "step_up_code_expired"
const StepUpErrorCodesStepUpLocked = "step_up_locked"

const ThreadWaitDefaultTimeoutSeconds = 30
const ThreadWaitDeadlineMarginSeconds = 15

var UnspecifiedHostnames = []string{"0.0.0.0", "::", "[::]"}

const WebhookHeadersSignature = "X-Webhook-Signature"
const WebhookHeadersEvent = "X-Webhook-Event"
const WebhookHeadersDelivery = "X-Webhook-Delivery"

const WebhookSignatureVersion = "v1"
const WebhookSignatureTimestamp = "t"
const WebhookSignaturePairSeparator = ","
const WebhookSignatureValueSeparator = "="
const WebhookSignatureToleranceSeconds = 300
