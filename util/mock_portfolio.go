package util

import "math/rand"

func RandomInt64(min, max int64) int64 {
	return rand.Int63n(max - min) + min
}