package security

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashInternal(t *testing.T) {
	ast := assert.New(t)
	h, actualMethod := hashInternal("pbkdf2:sha256", "iuAaxWvA", "123456")
	ast.Equal("pbkdf2:sha256:150000", actualMethod)
	ast.Equal("2f3831dc844683f87cd94786bb64c939d709f5a0518e0168bce150a19cc02f27", h)
}

func TestCheckPasswordHash(t *testing.T) {
	ast := assert.New(t)
	ret := CheckPasswordHash(
		"pbkdf2:sha256:150000$iuAaxWvA$2f3831dc844683f87cd94786bb64c939d709f5a0518e0168bce150a19cc02f27",
		"123456")
	ast.Equal(true, ret)
}
