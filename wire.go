package nextmail

import (
	"bytes"
	"encoding"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"nextmail.sh/internal/protocol"
)

const (
	bodyDepth            = 512
	instantLayout        = "2006-01-02T15:04:05.000Z"
	fixedNotationFloor   = 1e-6
	fixedNotationCeil    = 1e21
	fieldAttachments     = "attachments"
	fieldFileID          = "fileId"
	fieldURL             = "url"
	fieldFilename        = "filename"
	fieldContent         = "content"
	fieldContentType     = "contentType"
	replacementCharacter = "\uFFFD"
)

var (
	recipientFields      = []string{"to", "cc", "bcc"}
	linkAttachmentFields = []string{fieldFilename, fieldContentType}
)

var (
	timeType          = reflect.TypeOf(time.Time{})
	jsonMarshalerType = reflect.TypeOf((*json.Marshaler)(nil)).Elem()
	textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()
)

type Body = map[string]any

func instant(value time.Time) string {
	return value.UTC().Format(instantLayout)
}

func jsNumber(value float64) string {
	switch {
	case math.IsNaN(value):
		return "NaN"
	case math.IsInf(value, 1):
		return "Infinity"
	case math.IsInf(value, -1):
		return "-Infinity"
	case value == 0:
		return "0"
	}

	size := math.Abs(value)

	if size >= fixedNotationFloor && size < fixedNotationCeil {
		return strconv.FormatFloat(value, 'f', -1, 64)
	}

	text := strconv.FormatFloat(value, 'e', -1, 64)
	mantissa, exponent, _ := strings.Cut(text, "e")
	sign := exponent[:1]

	return mantissa + "e" + sign + strings.TrimLeft(exponent[1:], "0")
}

func queryText(value any) (string, bool) {
	switch typed := value.(type) {
	case nil:
		return "", false
	case string:
		return typed, typed != ""
	case bool:
		return strconv.FormatBool(typed), true
	case int:
		return strconv.Itoa(typed), true
	case int64:
		return strconv.FormatInt(typed, 10), true
	case float64:
		return jsNumber(typed), true
	case time.Time:
		return instant(typed), true
	case []string:
		joined := strings.Join(typed, protocol.ListSeparator)

		return joined, joined != ""
	}

	text := fmt.Sprint(value)

	return text, text != ""
}

func jsonText(value any) []byte {
	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(value); err != nil {
		return []byte("null")
	}

	return bytes.TrimRight(buffer.Bytes(), "\n")
}

func encodeBody(body any) ([]byte, error) {
	ready, err := jsonReady(reflect.ValueOf(body), 0)

	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)

	if err := encoder.Encode(ready); err != nil {
		return nil, argumentFailure(err, messageJSONEncode, err.Error())
	}

	return bytes.TrimRight(buffer.Bytes(), "\n"), nil
}

func jsonReady(value reflect.Value, depth int) (any, error) {
	if depth > bodyDepth {
		return nil, argumentError(messageBodyDepth, bodyDepth)
	}

	if !value.IsValid() {
		return nil, nil
	}

	kind := value.Kind()

	if kind == reflect.Interface || kind == reflect.Pointer {
		if value.IsNil() {
			return nil, nil
		}

		if kind == reflect.Pointer && value.Type().Elem() != timeType && marshalsItself(value.Type()) {
			return marshalled(value)
		}

		return jsonReady(value.Elem(), depth+1)
	}

	if value.Type() == timeType {
		return instant(value.Interface().(time.Time)), nil
	}

	if marshalsItself(value.Type()) {
		return marshalled(value)
	}

	switch kind {
	case reflect.String:
		return value.String(), nil
	case reflect.Bool:
		return value.Bool(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return finite(value.Float())
	case reflect.Map:
		return objectReady(value, depth)
	case reflect.Slice:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			return base64.StdEncoding.EncodeToString(value.Bytes()), nil
		}

		return listReady(value, depth)
	case reflect.Array:
		return listReady(value, depth)
	case reflect.Struct:
		return marshalled(value)
	}

	return nil, argumentError(messageBodyValueShape, value.Type().String())
}

func marshalsItself(kind reflect.Type) bool {
	return kind.Implements(jsonMarshalerType) || kind.Implements(textMarshalerType)
}

func marshalled(value reflect.Value) (any, error) {
	encoded, err := json.Marshal(value.Interface())

	if err != nil {
		return nil, argumentFailure(err, messageJSONEncode, err.Error())
	}

	return json.RawMessage(encoded), nil
}

func finite(value float64) (any, error) {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return nil, argumentError(messageBodyNumber, jsNumber(value))
	}

	if value == 0 {
		return 0, nil
	}

	return value, nil
}

func objectReady(value reflect.Value, depth int) (any, error) {
	if value.Type().Key().Kind() != reflect.String {
		return nil, argumentError(messageBodyValueShape, value.Type().String())
	}

	ready := make(map[string]any, value.Len())
	entries := value.MapRange()

	for entries.Next() {
		member, err := jsonReady(entries.Value(), depth+1)

		if err != nil {
			return nil, err
		}

		ready[entries.Key().String()] = member
	}

	return ready, nil
}

func listReady(value reflect.Value, depth int) (any, error) {
	ready := make([]any, 0, value.Len())

	for index := 0; index < value.Len(); index++ {
		member, err := jsonReady(value.Index(index), depth+1)

		if err != nil {
			return nil, err
		}

		ready = append(ready, member)
	}

	return ready, nil
}

func bodyOf(body Body) Body {
	if body == nil {
		return Body{}
	}

	return body
}

func listOf(value any) ([]any, bool) {
	if value == nil {
		return nil, false
	}

	reflected := reflect.ValueOf(value)

	if reflected.Kind() != reflect.Slice && reflected.Kind() != reflect.Array {
		return nil, false
	}

	if reflected.Type().Elem().Kind() == reflect.Uint8 {
		return nil, false
	}

	items := make([]any, 0, reflected.Len())

	for index := 0; index < reflected.Len(); index++ {
		items = append(items, reflected.Index(index).Interface())
	}

	return items, true
}

func mapOf(value any) (map[string]any, bool) {
	switch typed := value.(type) {
	case map[string]any:
		return typed, typed != nil
	case Object:
		return typed, typed != nil
	}

	if value == nil {
		return nil, false
	}

	reflected := reflect.ValueOf(value)

	if reflected.Kind() != reflect.Map || reflected.Type().Key().Kind() != reflect.String || reflected.IsNil() {
		return nil, false
	}

	entries := make(map[string]any, reflected.Len())
	members := reflected.MapRange()

	for members.Next() {
		entries[members.Key().String()] = members.Value().Interface()
	}

	return entries, true
}

func many(value any) any {
	if value == nil {
		return nil
	}

	if _, ok := listOf(value); ok {
		return value
	}

	return []any{value}
}

func recipientsToWire(body Body) Body {
	wire := make(Body, len(body))

	for key, value := range body {
		wire[key] = value
	}

	for _, field := range recipientFields {
		if value, ok := wire[field]; ok {
			wire[field] = many(value)
		}
	}

	return wire
}

func emailToWire(body Body) (Body, error) {
	return attachmentsToWire(recipientsToWire(body))
}

func attachmentsToWire(body Body) (Body, error) {
	attachments, present := body[fieldAttachments]

	if !present || attachments == nil {
		return body, nil
	}

	items, ok := listOf(attachments)

	if !ok {
		return nil, argumentError(messageAttachmentsShape)
	}

	prepared := make([]any, 0, len(items))

	for _, item := range items {
		attachment, err := attachmentToWire(item)

		if err != nil {
			return nil, err
		}

		prepared = append(prepared, attachment)
	}

	wire := maps.Clone(body)
	wire[fieldAttachments] = prepared

	return wire, nil
}

func attachmentToWire(value any) (Body, error) {
	attachment, ok := mapOf(value)

	if !ok {
		return nil, argumentError(messageAttachmentShape)
	}

	if fileID, stored := attachment[fieldFileID]; stored {
		return Body{fieldFileID: fileID}, nil
	}

	if link, linked := attachment[fieldURL]; linked {
		wire := Body{fieldURL: link}

		for _, field := range linkAttachmentFields {
			if given, named := attachment[field]; named && given != nil {
				wire[field] = given
			}
		}

		return wire, nil
	}

	content, err := attachmentContent(attachment[fieldContent])

	if err != nil {
		return nil, err
	}

	wire := Body{fieldFilename: attachment[fieldFilename], fieldContent: content}

	if contentType, typed := attachment[fieldContentType]; typed && contentType != nil {
		wire[fieldContentType] = contentType
	}

	return wire, nil
}

func attachmentContent(value any) (string, error) {
	switch typed := value.(type) {
	case string:
		return typed, nil
	case []byte:
		return base64.StdEncoding.EncodeToString(typed), nil
	case io.Reader:
		data, err := io.ReadAll(typed)

		if err != nil {
			return "", argumentFailure(err, messageAttachmentUnreadable, err.Error())
		}

		return base64.StdEncoding.EncodeToString(data), nil
	}

	return "", argumentError(messageAttachmentContent)
}
