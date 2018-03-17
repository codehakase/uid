package uid

import (
	"math/rand"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())

const (
	characterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// chacterBits represents a character index
	characterBits        = 6
	characterIndicesMax  = 63 / characterBits
	characterIndicesMask = 1<<characterBits - 1
)

// New returns a random string of length n
func New(n int) string {
	return NewFromBytes(n, characterBytes)
}

// NewFromBytes returns random string of length n from characterBytes
func NewFromBytes(n int, char string) string {
	bytes := make([]byte, n)

	for i, c, r := n-1, seed.Int63(), characterIndicesMax; i >= 0; {
		if r == 0 {
			c, r = seed.Int63(), characterIndicesMax
		}
		if index := int(c & characterIndicesMask); index < len(char) {
			bytes[i] = char[index]
			i--
		}
		c >>= characterBits
		r--
	}

	return string(bytes)
}
