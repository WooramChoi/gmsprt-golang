package utils

import (
	"strings"
	"unicode"
)

func Substr(str string, length int) string {
	var rslt []rune
	for idx, val := range str {
		if idx >= length {
			break
		}
		rslt = append(rslt, val)
	}
	return string(rslt)
}

func ToSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('_')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}
