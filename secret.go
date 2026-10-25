package nextmail

import (
	"fmt"
	"log/slog"
)

type secret struct {
	reveal func() string
}

func newSecret(value string) secret {
	if value == "" {
		return secret{}
	}

	return secret{reveal: func() string { return value }}
}

func (s secret) present() bool {
	return s.reveal != nil
}

func (s secret) value() string {
	if s.reveal == nil {
		return ""
	}

	return s.reveal()
}

func (s secret) String() string {
	if s.reveal == nil {
		return ""
	}

	return redacted
}

func (s secret) GoString() string {
	return s.String()
}

func (s secret) Format(state fmt.State, _ rune) {
	_, _ = fmt.Fprint(state, s.String())
}

func (s secret) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s secret) LogValue() slog.Value {
	return slog.StringValue(s.String())
}
