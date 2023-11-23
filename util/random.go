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
func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
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

func RandomYear() int32 {
	return RandomInt32(1950, 2023)
}
