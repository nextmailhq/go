package nextmail

import (
	"math"
	"net/netip"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"nextmail.sh/internal/protocol"
)

const (
	upperHex               = "0123456789ABCDEF"
	componentUnreserved    = "-_.!~*'()"
	formUnreserved         = "*-._"
	maxIPv4Parts           = 4
	schemeHTTP             = "http"
	schemeHTTPS            = "https"
	defaultHTTPPort        = "80"
	defaultHTTPSPort       = "443"
	trailingSlash          = "/"
	maxIPv4PartDigits      = 15
	ipv4Base               = 256
	loopbackIPv4Expression = `^127\.\d{1,3}\.\d{1,3}\.\d{1,3}$`
)

var (
	loopbackIPv4Pattern      = regexp.MustCompile(loopbackIPv4Expression)
	endpointSchemePattern    = regexp.MustCompile(`(?i)^[a-z][a-z0-9+.-]*://`)
	userInfoPattern          = regexp.MustCompile(`(?i)^([a-z][a-z0-9+.-]*://)[^/?#]*@`)
	pathParamPattern         = regexp.MustCompile(`:([A-Za-z0-9_]+)(?:\{[^}]*\})?`)
	optionalPathParamPattern = regexp.MustCompile(`/:([A-Za-z0-9_]+)\?`)
	dotSegmentPattern        = regexp.MustCompile(`^\.+$`)
	ipv4DecimalPattern       = regexp.MustCompile(`^(?:0|[1-9]\d*)$`)
	ipv4OctalPattern         = regexp.MustCompile(`^0[0-7]+$`)
	ipv4HexPattern           = regexp.MustCompile(`(?i)^0x[0-9a-f]*$`)
)

func normaliseEndpoint(value string) string {
	text := strings.TrimSpace(value)
	trimmed := strings.TrimRight(text, trailingSlash)

	if endpointSchemePattern.MatchString(text) {
		return trimmed
	}

	if trimmed == "" {
		return ""
	}

	if isLocalHost(hostnameOf(trimmed)) {
		return schemeHTTP + "://" + trimmed
	}

	return schemeHTTPS + "://" + trimmed
}

func hostnameOf(authority string) string {
	parsed, err := url.Parse(schemeHTTP + "://" + authority)

	if err != nil {
		return ""
	}

	return canonicalHost(parsed.Hostname())
}

func parseHTTP(text string) (*url.URL, bool) {
	if hasForbidden(text, ' ') {
		return nil, false
	}

	parsed, err := url.Parse(text)

	if err != nil || parsed.Opaque != "" || parsed.Hostname() == "" {
		return nil, false
	}

	if parsed.Scheme != schemeHTTP && parsed.Scheme != schemeHTTPS {
		return nil, false
	}

	if port := parsed.Port(); port != "" {
		if number, err := strconv.Atoi(port); err != nil || number > math.MaxUint16 {
			return nil, false
		}
	}

	return parsed, true
}

func hasForbidden(text string, ceiling rune) bool {
	for _, character := range text {
		if character <= ceiling || character == 0x7f {
			return true
		}
	}

	return false
}

func validBaseURL(value string) (string, error) {
	text := strings.TrimRight(strings.TrimSpace(value), trailingSlash)
	parsed, ok := parseHTTP(text)

	if !ok {
		shown := userInfoPattern.ReplaceAllString(strings.TrimSpace(value), "$1")

		return "", argumentError("%s %s", quoted(shown), messageBaseURLShape)
	}

	if parsed.User != nil {
		return "", argumentError(messageBaseURLUserInfo)
	}

	host := canonicalHost(parsed.Hostname())

	if isUnspecifiedHost(host) {
		return "", argumentError(messageUnspecifiedHost, quoted(value), host, loopbackSuggestion(parsed))
	}

	return text, nil
}

func loopbackSuggestion(parsed *url.URL) string {
	host := protocol.LoopbackIpv4Address

	if strings.Contains(parsed.Hostname(), ":") {
		host = protocol.LoopbackIpv6Address
	}

	if port := parsed.Port(); port != "" {
		host += ":" + port
	}

	return strings.TrimRight(parsed.Scheme+"://"+host+parsed.EscapedPath(), trailingSlash)
}

func isSecureOrigin(baseURL string) bool {
	parsed, ok := parseHTTP(baseURL)

	if !ok {
		return false
	}

	return parsed.Scheme == schemeHTTPS || isLoopbackHost(canonicalHost(parsed.Hostname()))
}

func originOf(text string) (string, bool) {
	parsed, ok := parseHTTP(text)

	if !ok {
		return "", false
	}

	port := parsed.Port()

	if port == "" {
		port = defaultHTTPPort

		if parsed.Scheme == schemeHTTPS {
			port = defaultHTTPSPort
		}
	}

	return parsed.Scheme + "://" + canonicalHost(parsed.Hostname()) + ":" + port, true
}

func apiURL(baseURL, path string) (string, error) {
	if !strings.HasPrefix(path, "/") || strings.HasPrefix(path, "//") || strings.HasPrefix(path, "/\\") || hasForbidden(path, 0x1f) {
		return "", argumentError("%s %s", quoted(path), messagePathShape)
	}

	full := strings.TrimRight(baseURL, trailingSlash) + path
	origin, ok := originOf(full)
	base, _ := originOf(baseURL)

	if !ok || origin != base {
		return "", argumentError("%s %s", quoted(path), messageOffOrigin)
	}

	return full, nil
}

func buildURL(baseURL, path string, parameters values) (string, error) {
	full, err := apiURL(baseURL, path)

	if err != nil {
		return "", err
	}

	pairs := make([]string, 0, len(parameters))

	for _, entry := range parameters {
		text, ok := queryText(entry.value)

		if !ok {
			continue
		}

		pairs = append(pairs, formEncode(entry.key)+"="+formEncode(text))
	}

	if len(pairs) == 0 {
		return full, nil
	}

	return full + "?" + strings.Join(pairs, "&"), nil
}

func percentEncode(value, unreserved string, plusForSpace bool) string {
	var encoded strings.Builder

	for index := 0; index < len(value); index++ {
		character := value[index]

		switch {
		case character >= 'A' && character <= 'Z', character >= 'a' && character <= 'z', character >= '0' && character <= '9':
			encoded.WriteByte(character)
		case strings.IndexByte(unreserved, character) >= 0:
			encoded.WriteByte(character)
		case character == ' ' && plusForSpace:
			encoded.WriteByte('+')
		default:
			encoded.WriteByte('%')
			encoded.WriteByte(upperHex[character>>4])
			encoded.WriteByte(upperHex[character&0x0f])
		}
	}

	return encoded.String()
}

func formEncode(value string) string {
	return percentEncode(strings.ToValidUTF8(value, replacementCharacter), formUnreserved, true)
}

func encodeURIComponent(value string) string {
	return percentEncode(value, componentUnreserved, false)
}

func isLoopbackHost(host string) bool {
	return inHosts(protocol.LoopbackHostnames, host) || loopbackIPv4Pattern.MatchString(host)
}

func isLocalHost(host string) bool {
	return inHosts(protocol.LocalHosts, host) || loopbackIPv4Pattern.MatchString(host)
}

func isUnspecifiedHost(host string) bool {
	return inHosts(protocol.UnspecifiedHostnames, host)
}

func inHosts(hosts []string, host string) bool {
	for _, candidate := range hosts {
		if strings.Trim(candidate, "[]") == host {
			return true
		}
	}

	return false
}

func canonicalHost(host string) string {
	text := strings.ToLower(strings.Trim(host, "[]"))

	if text == "" {
		return text
	}

	if strings.Contains(text, ":") {
		address, err := netip.ParseAddr(text)

		if err != nil {
			return text
		}

		return address.WithZone("").String()
	}

	parts := strings.Split(text, ".")

	if len(parts) > 1 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	if len(parts) > maxIPv4Parts {
		return text
	}

	numbers := make([]float64, 0, len(parts))

	for _, part := range parts {
		number, ok := ipv4Number(part)

		if !ok {
			return text
		}

		numbers = append(numbers, number)
	}

	address, ok := ipv4Address(numbers)

	if !ok {
		return text
	}

	return address
}

func ipv4Number(part string) (float64, bool) {
	switch {
	case ipv4HexPattern.MatchString(part):
		if len(part) == 2 {
			return 0, true
		}

		return parsedNumber(part[2:], 16)
	case ipv4OctalPattern.MatchString(part):
		return parsedNumber(part, 8)
	case ipv4DecimalPattern.MatchString(part):
		if len(part) > maxIPv4PartDigits {
			return math.Inf(1), true
		}

		return parsedNumber(part, 10)
	}

	return 0, false
}

func parsedNumber(digits string, base int) (float64, bool) {
	number, err := strconv.ParseUint(digits, base, 64)

	if err != nil {
		return math.Inf(1), true
	}

	return float64(number), true
}

func ipv4Address(numbers []float64) (string, bool) {
	count := len(numbers)

	if count == 0 {
		return "", false
	}

	last := numbers[count-1]
	leading := numbers[:count-1]

	for _, number := range leading {
		if number > 255 {
			return "", false
		}
	}

	if last >= math.Pow(ipv4Base, float64(maxIPv4Parts+1-count)) {
		return "", false
	}

	value := uint32(last)

	for index, number := range leading {
		value += uint32(number) << (8 * (3 - index))
	}

	return strconv.Itoa(int(value>>24&255)) + "." + strconv.Itoa(int(value>>16&255)) + "." + strconv.Itoa(int(value>>8&255)) + "." + strconv.Itoa(int(value&255)), true
}

func fillPath(pattern string, params []any) (string, error) {
	values := make(map[string]any, len(params)/2)

	for index := 0; index+1 < len(params); index += 2 {
		if name, ok := params[index].(string); ok {
			values[name] = params[index+1]
		}
	}

	texts := make(map[string]string, len(values))

	for _, match := range pathParamPattern.FindAllStringSubmatch(pattern, -1) {
		text, err := segmentText(values[match[1]])

		if err != nil {
			return "", err
		}

		texts[match[1]] = text
	}

	optional := optionalPathParamPattern.ReplaceAllStringFunc(pattern, func(found string) string {
		name := optionalPathParamPattern.FindStringSubmatch(found)[1]

		return "/" + encodeURIComponent(texts[name])
	})

	return pathParamPattern.ReplaceAllStringFunc(optional, func(found string) string {
		name := pathParamPattern.FindStringSubmatch(found)[1]

		return encodeURIComponent(texts[name])
	}), nil
}

func segmentText(value any) (string, error) {
	var text string

	switch typed := value.(type) {
	case nil:
		return "", argumentError(messageEmptySegment)
	case string:
		text = typed
	case int:
		text = strconv.Itoa(typed)
	case int64:
		text = strconv.FormatInt(typed, 10)
	default:
		return "", argumentError(messageSegmentShape, value)
	}

	if text == "" {
		return "", argumentError(messageEmptySegment)
	}

	if !utf8.ValidString(text) {
		return "", argumentError(messageSegmentEncoding)
	}

	if dotSegmentPattern.MatchString(text) {
		return "", argumentError("%s %s", quoted(text), messageDotSegment)
	}

	return text, nil
}
