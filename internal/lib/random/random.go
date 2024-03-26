package random

import (
	"math/rand"
	"time"
)

// NewRandomString generates random string with given size.
func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	result := make([]rune, size)
	for i := range result {
		result[i] = chars[rnd.Intn(len(chars))]
	}

	return string(result)
}
