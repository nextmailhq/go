package nextmail

import (
	"slices"
	"strings"
	"sync"

	"nextmail.sh/internal/protocol"
)

const languageSubtagSeparators = "-_"

type Language struct {
	Code   string
	Label  string
	Native string
	Flag   string
	RTL    bool
}

type languageIndex struct {
	byCode   map[string]Language
	byLabel  map[string]Language
	byNative map[string]Language
	aliases  map[string]string
}

var languageTable = sync.OnceValue(func() *languageIndex {
	index := &languageIndex{
		byCode:   make(map[string]Language, len(languages)),
		byLabel:  make(map[string]Language, len(languages)),
		byNative: make(map[string]Language, len(languages)),
		aliases:  make(map[string]string, len(protocol.LanguageAliases)),
	}

	for _, language := range languages {
		index.byCode[languageKey(language.Code)] = language
		index.byLabel[languageKey(language.Label)] = language
		index.byNative[languageKey(language.Native)] = language
	}

	for _, alias := range protocol.LanguageAliases {
		if len(alias) == 2 {
			index.aliases[languageKey(alias[0])] = languageKey(alias[1])
		}
	}

	return index
})

func languageKey(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func Languages() []Language {
	return slices.Clone(languages)
}

func LanguageByCode(code string) (Language, bool) {
	language, found := languageTable().byCode[languageKey(code)]

	return language, found
}

func ResolveLanguage(input string) (Language, bool) {
	key := languageKey(input)

	if key == "" {
		return Language{}, false
	}

	table := languageTable()

	if aliased, known := table.aliases[key]; known {
		language, found := table.byCode[aliased]

		return language, found
	}

	for _, names := range []map[string]Language{table.byCode, table.byLabel, table.byNative} {
		if language, found := names[key]; found {
			return language, true
		}
	}

	base := key

	if at := strings.IndexAny(key, languageSubtagSeparators); at >= 0 {
		base = key[:at]
	}

	if base == "" || base == key {
		return Language{}, false
	}

	if aliased, known := table.aliases[base]; known {
		base = aliased
	}

	language, found := table.byCode[base]

	return language, found
}

func IsRTLLanguage(code string) bool {
	language, found := ResolveLanguage(code)

	return found && language.RTL
}
