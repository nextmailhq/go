package nextmail

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nextmail.sh/internal/protocol"
)

type verifyConfig struct {
	tolerance time.Duration
	now       func() time.Time
	err       error
}

type VerifyOption func(config *verifyConfig)

func WithTolerance(tolerance time.Duration) VerifyOption {
	return func(config *verifyConfig) {
		if tolerance < 0 {
			config.err = argumentError(messageWebhookToleranceSign)

			return
		}

		config.tolerance = tolerance
	}
}

func VerifyWebhookSignature(payload []byte, headers http.Header, signingSecret string, opts ...VerifyOption) (Object, error) {
	config := &verifyConfig{tolerance: protocol.WebhookSignatureToleranceSeconds * time.Second, now: time.Now}

	for _, option := range opts {
		if option != nil {
			option(config)
		}
	}

	if config.err != nil {
		return nil, config.err
	}

	if strings.TrimSpace(signingSecret) == "" {
		return nil, argumentError(messageWebhookSecret)
	}

	signature := headerNamed(headers, protocol.WebhookHeadersSignature)

	if signature == "" {
		return nil, signatureError(nil, messageWebhookMissingHeader)
	}

	timestamp, candidates, ok := parseSignature(signature)

	if !ok {
		return nil, signatureError(nil, messageWebhookMalformed)
	}

	tolerance := config.tolerance.Seconds()
	drift := math.Abs(math.Floor(float64(config.now().UnixMilli())/1000) - timestamp)

	if tolerance > 0 && drift > tolerance {
		return nil, signatureError(nil, messageWebhookTolerance, jsNumber(drift), jsNumber(tolerance))
	}

	mac := hmac.New(sha256.New, []byte(signingSecret))
	mac.Write([]byte(jsNumber(timestamp) + "."))
	mac.Write(payload)

	expected := []byte(hex.EncodeToString(mac.Sum(nil)))
	matched := 0

	for _, candidate := range candidates {
		matched |= subtle.ConstantTimeCompare(expected, []byte(candidate))
	}

	if matched != 1 {
		return nil, signatureError(nil, messageWebhookMismatch)
	}

	decoded, valid := decodeJSON([]byte(textOf(payload)))
	event := objectOf(decoded)

	if !valid || event == nil {
		return nil, signatureError(nil, messageWebhookPayloadShape)
	}

	return event, nil
}

func headerNamed(headers http.Header, name string) string {
	for key, values := range headers {
		if strings.EqualFold(key, name) && len(values) > 0 {
			return values[0]
		}
	}

	return ""
}

func parseSignature(header string) (float64, []string, bool) {
	timestamp := math.NaN()
	signatures := []string{}

	for _, entry := range strings.Split(header, protocol.WebhookSignaturePairSeparator) {
		at := strings.Index(entry, protocol.WebhookSignatureValueSeparator)

		if at < 1 {
			continue
		}

		key := strings.TrimSpace(entry[:at])
		value := strings.TrimSpace(entry[at+1:])

		if key == protocol.WebhookSignatureTimestamp {
			timestamp = numberOf(value)
		}

		if key == protocol.WebhookSignatureVersion && value != "" {
			signatures = append(signatures, value)
		}
	}

	if math.IsNaN(timestamp) || math.IsInf(timestamp, 0) || len(signatures) == 0 {
		return 0, nil, false
	}

	return timestamp, signatures, true
}

func numberOf(value string) float64 {
	if value == "" {
		return 0
	}

	if !decimalPattern.MatchString(value) {
		return math.NaN()
	}

	number, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return math.NaN()
	}

	return number
}
