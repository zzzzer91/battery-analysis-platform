package security

import (
	"crypto/rand"
	"testing"
)

func BenchmarkGenerateRandomString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = GenerateRandomString(32)
	}
}

func BenchmarkRandom(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 32)
		_, _ = rand.Read(buf)
	}
}
