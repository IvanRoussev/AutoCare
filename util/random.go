package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min and max
func RandomInt64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString() string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < 6; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomYear() int64 {
	return RandomInt64(1950, 2023)
}
