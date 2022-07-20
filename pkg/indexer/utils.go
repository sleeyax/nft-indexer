package indexer

import "strings"

// Normalize normalizes a blockchain address (trim + lower case).
//
// Addresses should be normalized before storing them in a database.
// That way can rest assured the correct result is always returned when comparing two different addresses, for example.
func Normalize(address string) string {
	return strings.ToLower(strings.TrimSpace(address))
}
