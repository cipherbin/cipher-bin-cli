package randstring

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// New creates a random string with caller supplied length
func New(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}
	return string(b)
}
