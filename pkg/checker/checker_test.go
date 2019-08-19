package checker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChecker(t *testing.T) {
	ast := assert.New(t)
	ret := ReUserNameOrPassword.MatchString("xiaoming")
	expected := true
	ast.Equal(expected, ret)
}
