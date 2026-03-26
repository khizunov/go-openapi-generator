package goresty

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English)

var wellKnownAcronyms = map[string]string{
	"id":    "ID",
	"url":   "URL",
	"uri":   "URI",
	"http":  "HTTP",
	"https": "HTTPS",
	"api":   "API",
	"uid":   "UID",
	"uuid":  "UUID",
	"ip":    "IP",
	"sql":   "SQL",
}

// ToModelName converts a spec name to an exported Go type name (PascalCase).
func ToModelName(name string) string {
	if name == "" {
		return ""
	}
	words := splitWords(name)
	for i, w := range words {
		lower := strings.ToLower(w)
		if acro, ok := wellKnownAcronyms[lower]; ok {
			words[i] = acro
		} else {
			words[i] = titleCaser.String(lower)
		}
	}
	return strings.Join(words, "")
}

// ToOperationName converts an operationId to an exported Go function name.
func ToOperationName(name string) string {
	return ToModelName(name)
}

// ToPropertyName converts a property name to an exported Go field name.
func ToPropertyName(name string) string {
	return ToModelName(name)
}

// ToParamName converts a parameter name to an unexported Go variable name (camelCase).
func ToParamName(name string) string {
	if name == "" {
		return ""
	}
	name = strings.ReplaceAll(name, "-", "_")
	words := splitWords(name)
	for i, w := range words {
		lower := strings.ToLower(w)
		if i == 0 {
			words[i] = strings.ToLower(w)
			continue
		}
		if acro, ok := wellKnownAcronyms[lower]; ok {
			words[i] = acro
		} else {
			words[i] = titleCaser.String(lower)
		}
	}
	if len(words) > 0 {
		words[0] = strings.ToLower(words[0])
	}
	return strings.Join(words, "")
}

func splitWords(s string) []string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})
	var words []string
	for _, part := range parts {
		words = append(words, splitCamel(part)...)
	}
	return words
}

func splitCamel(s string) []string {
	var words []string
	var current strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if i > 0 && unicode.IsUpper(r) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}
	return words
}
