package nextmail

import (
	"math"
	"time"
)

const (
	fieldData       = "data"
	fieldHasMore    = "hasMore"
	fieldNextCursor = "nextCursor"
	fieldID         = "id"
)

type Object map[string]any

func (o Object) ID() string {
	return o.String(fieldID)
}

func (o Object) Has(key string) bool {
	value, present := o[key]

	return present && value != nil
}

func (o Object) String(key string) string {
	text, _ := o[key].(string)

	return text
}

func (o Object) Float(key string) float64 {
	switch typed := o[key].(type) {
	case float64:
		return typed
	case int:
		return float64(typed)
	case int64:
		return float64(typed)
	}

	return 0
}

func (o Object) Int(key string) int {
	number := o.Float(key)

	if math.IsNaN(number) || math.IsInf(number, 0) {
		return 0
	}

	return int(number)
}

func (o Object) Bool(key string) bool {
	flag, _ := o[key].(bool)

	return flag
}

func (o Object) Time(key string) time.Time {
	text := o.String(key)

	if text == "" {
		return time.Time{}
	}

	parsed, err := time.Parse(time.RFC3339Nano, text)

	if err != nil {
		return time.Time{}
	}

	return parsed
}

func (o Object) Object(key string) Object {
	return objectOf(o[key])
}

func (o Object) Objects(key string) []Object {
	return objectsOf(o[key])
}

func (o Object) Strings(key string) []string {
	items, ok := o[key].([]any)

	if !ok {
		return nil
	}

	texts := make([]string, 0, len(items))

	for _, item := range items {
		if text, ok := item.(string); ok {
			texts = append(texts, text)
		}
	}

	return texts
}

func objectOf(value any) Object {
	switch typed := value.(type) {
	case Object:
		return typed
	case map[string]any:
		return typed
	}

	return nil
}

func objectsOf(value any) []Object {
	items, ok := value.([]any)

	if !ok {
		return []Object{}
	}

	objects := make([]Object, 0, len(items))

	for _, item := range items {
		if object := objectOf(item); object != nil {
			objects = append(objects, object)
		}
	}

	return objects
}

func cursorOf(value any) string {
	text, _ := value.(string)

	return text
}

func hasMoreOf(envelope Object, nextCursor string) bool {
	if flag, ok := envelope[fieldHasMore].(bool); ok {
		return flag
	}

	return nextCursor != ""
}

type Page struct {
	Items      []Object
	HasMore    bool
	NextCursor string
}

type AddressBook struct {
	Unrestricted bool
	Addresses    []Object
	Domains      []Object
}

type AddressBookPage struct {
	Unrestricted bool
	Addresses    []Object
	Domains      []Object
	HasMore      bool
	NextCursor   string
}

type PeoplePage struct {
	Items      []Object
	HasMore    bool
	NextCursor string
	Seen       bool
}

type TempMessagesPage struct {
	Items      []Object
	HasMore    bool
	NextCursor string
	ExpiresAt  string
}

type BatchResult struct {
	Items  []Object
	Sent   int
	Failed int
}

type TemplateSends struct {
	Items    []Object
	Total    int
	Page     int
	PageSize int
}
