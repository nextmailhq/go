package nextmail

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"time"

	"nextmail.sh/internal/protocol"
)

var (
	ErrAuthentication   = errors.New("nextmail: the credential was refused")
	ErrPermission       = errors.New("nextmail: the credential may not do this")
	ErrScopeMissing     = errors.New("nextmail: the credential lacks a scope")
	ErrStepUpRequired   = errors.New("nextmail: a verification code is required")
	ErrInvalidRequest   = errors.New("nextmail: the request was not valid")
	ErrValidation       = errors.New("nextmail: a field failed validation")
	ErrNotFound         = errors.New("nextmail: not found")
	ErrConflict         = errors.New("nextmail: the request conflicts with the current state")
	ErrRateLimited      = errors.New("nextmail: rate limited")
	ErrServer           = errors.New("nextmail: the API failed")
	ErrRetryable        = errors.New("nextmail: the request can be tried again")
	ErrNetwork          = errors.New("nextmail: no response arrived")
	ErrTimeout          = errors.New("nextmail: the request timed out")
	ErrInvalidArgument  = errors.New("nextmail: an argument was refused before anything was sent")
	ErrWebhookSignature = errors.New("nextmail: the webhook signature was refused")
)

type errorKind int

const (
	kindAPI errorKind = iota
	kindNetwork
	kindArgument
	kindSignature
)

type FieldProblem struct {
	Key   string
	Error string
}

type Error struct {
	Message    string
	Status     int
	Type       string
	Code       string
	Param      string
	DocURL     string
	RequestID  string
	RetryAfter time.Duration
	Fields     []FieldProblem
	Body       any

	kind    errorKind
	timeout bool
	cause   error
}

func (e *Error) Error() string {
	if e == nil {
		return "nextmail: <nil>"
	}

	if e.kind != kindAPI {
		return "nextmail: " + e.Message
	}

	detail := strconv.Itoa(e.Status)

	if e.Code != "" {
		detail += " " + e.Code
	}

	if e.RequestID != "" {
		detail += ", request " + e.RequestID
	}

	return "nextmail: " + e.Message + " (" + detail + ")"
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.cause
}

func (e *Error) Is(target error) bool {
	if e == nil {
		return false
	}

	switch target {
	case ErrAuthentication:
		return e.isType(protocol.ErrorTypesAuthenticationError)
	case ErrPermission:
		return e.isType(protocol.ErrorTypesPermissionError)
	case ErrScopeMissing:
		return e.kind == kindAPI && e.Code == protocol.ErrorCodesInsufficientScope
	case ErrStepUpRequired:
		return e.kind == kindAPI && e.Code == protocol.StepUpErrorCodesStepUpRequired
	case ErrInvalidRequest:
		return e.isType(protocol.ErrorTypesInvalidRequestError)
	case ErrValidation:
		return e.isType(protocol.ErrorTypesValidationError)
	case ErrNotFound:
		return e.isType(protocol.ErrorTypesNotFoundError)
	case ErrConflict:
		return e.isType(protocol.ErrorTypesConflictError)
	case ErrRateLimited:
		return e.isType(protocol.ErrorTypesRateLimitError)
	case ErrServer:
		return e.kind == kindAPI && e.Status >= 500
	case ErrRetryable:
		return e.Retryable()
	case ErrNetwork:
		return e.kind == kindNetwork
	case ErrTimeout:
		return e.Timeout()
	case ErrInvalidArgument:
		return e.kind == kindArgument
	case ErrWebhookSignature:
		return e.kind == kindSignature
	}

	return false
}

func (e *Error) Retryable() bool {
	return e != nil && e.kind == kindAPI && slices.Contains(protocol.RetryStatuses, e.Status)
}

func (e *Error) Timeout() bool {
	return e != nil && e.kind == kindNetwork && e.timeout
}

func (e *Error) isType(name string) bool {
	return e.kind == kindAPI && e.Type == name
}

func argumentError(format string, values ...any) *Error {
	return &Error{Message: fmt.Sprintf(format, values...), kind: kindArgument}
}

func argumentFailure(cause error, format string, values ...any) *Error {
	return &Error{Message: fmt.Sprintf(format, values...), kind: kindArgument, cause: cause}
}

func networkError(cause error, timeout bool, format string, values ...any) *Error {
	return &Error{Message: fmt.Sprintf(format, values...), kind: kindNetwork, timeout: timeout, cause: cause}
}

func signatureError(cause error, format string, values ...any) *Error {
	return &Error{Message: fmt.Sprintf(format, values...), kind: kindSignature, cause: cause}
}

func quoted(value string) string {
	return string(jsonText(value))
}
