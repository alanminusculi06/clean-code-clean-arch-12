package shared

import (
	"strings"
)

func IsEmpty(input string) bool {
	return input == ""
}

func ClearString(value string, charsToRemove ...string) string {
	for _, charToRemove := range charsToRemove {
		value = strings.ReplaceAll(value, charToRemove, "")
	}
	return value
}
