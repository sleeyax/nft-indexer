package utils

import "math/rand"

func RandomItem[T int | string](items []T) T {
	return items[rand.Intn(len(items))]
}