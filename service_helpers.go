package nextmail

import "strings"

func pick(body Body, keys ...string) Body {
	picked := make(Body, len(keys))

	for _, key := range keys {
		if value, present := body[key]; present {
			picked[key] = value
		}
	}

	return picked
}

func optionsBody(options *callOptions, names ...string) Body {
	body := make(Body, len(names))

	for _, name := range names {
		if options.has(name) {
			body[name] = options.value(name)
		}
	}

	return body
}

func nullable(text *string) any {
	if text == nil {
		return nil
	}

	return *text
}

func uploadName(filename string) (string, error) {
	name := strings.TrimSpace(filename)

	if name == "" {
		return "", argumentError(messageFilenameRequired)
	}

	return name, nil
}
