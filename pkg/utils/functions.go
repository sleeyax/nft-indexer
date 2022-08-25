package utils

import "math/rand"

func RandomItem[T int | string](items []T) T {
	return items[rand.Intn(len(items))]
}

// OrString returns the first string that isn't a zero value.
// Returns a string zero value ("") if no specified string matches the predicate.
func OrString(strings ...string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}

	return ""
}

// Ternary operator-like replacement.
func Ternary[T int | string](predicate bool, left T, right T) T {
	if predicate {
		return left
	} else {
		return right
	}
}
