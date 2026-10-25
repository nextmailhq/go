package nextmail

import (
	"strings"
	"unicode/utf16"

	"nextmail.sh/internal/protocol"
)

func IsAPIKey(value string) bool {
	return strings.HasPrefix(value, protocol.APIKeyPrefixesLive) || strings.HasPrefix(value, protocol.APIKeyPrefixesTest)
}

func IsAccessToken(value string) bool {
	if value == "" || strings.HasPrefix(value, protocol.AccessTokenRulesReservedPrefix) {
		return false
	}

	return len(utf16.Encode([]rune(value))) <= protocol.AccessTokenRulesMaxLength
}

func IsSealed(message Object) bool {
	format := message.Object("encryption").String("format")

	for _, sealed := range protocol.SealedEncryptionFormats {
		if format == sealed {
			return true
		}
	}

	return false
}

func assertAPIKey(apiKey, subject string) error {
	if apiKey == "" {
		return argumentError(messageAPIKeyRequired)
	}

	if !IsAPIKey(apiKey) {
		return argumentError("%s %s", subject, messageAPIKeyShape)
	}

	return nil
}

func assertCredential(apiKey, accessToken secret, withTokenFunc, required bool) error {
	withKey := apiKey.present()
	withToken := accessToken.present() || withTokenFunc

	switch {
	case withKey && withToken:
		return argumentError(messageCredentialConflict)
	case withKey:
		return assertAPIKey(apiKey.value(), messageAPIKeySubject)
	case withTokenFunc:
		return nil
	case withToken:
		if !IsAccessToken(accessToken.value()) {
			return argumentError("%s %s", messageAccessTokenSubject, messageAccessTokenShape)
		}
	case required:
		return argumentError(messageCredentialRequired)
	}

	return nil
}
