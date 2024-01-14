package randstring

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// New creates a cryptographically secure random string with the specified length.
func New(length int) (string, error) {
	b := make([]byte, length)
	for i := range b {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[charIndex.Int64()]
	}
	return string(b), nil
}
