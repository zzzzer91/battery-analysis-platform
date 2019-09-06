package security

import (
	"math/rand"
	"time"
)

func GenerateSalt(length int) string {
	if length <= 0 {
		panic("Salt length must be positive")
	}
	saltChars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	saltCharsLen := len(saltChars)
	result := make([]byte, 0, saltCharsLen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, saltChars[r.Intn(saltCharsLen)])
	}
	return string(result)
}
