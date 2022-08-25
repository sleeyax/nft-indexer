package utils

import (
	"regexp"
	"strings"
)

var nonSearchFriendlyCharactersRegex = regexp.MustCompile(`[\s-_]`)

// ToSearchFriendly converts the specified slug to a search friendly string by removing special characters that would otherwise only appear in the URL.
func ToSearchFriendly(str string) string {
	friendly := nonSearchFriendlyCharactersRegex.ReplaceAllString(str, "")
	return strings.ToLower(friendly)
}
