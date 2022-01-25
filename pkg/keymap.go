package recon

import "strings"

// KeyMapper is any function that transforms the given key into the standard key
type KeyMapper func(string) string

func DefaultMapper(source string) string {
	return source
}

func CamelCaseMapper(source string) string {
	if len(source) > 1 {
		return strings.ToUpper(source[0:1]) + source[1:]
	} else if len(source) == 1 {
		return strings.ToUpper(source[0:1])
	} else {
		return source
	}
}
