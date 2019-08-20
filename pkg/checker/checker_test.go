package checker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReUserNameOrPassword(t *testing.T) {
	ast := assert.New(t)
	testCases := []string{"xiaoming", "XIAOMING", "123456", "abcd", "xiaoming12345678"}
	expected := []bool{true, true, true, false, false}
	for i := range testCases {
		ret := ReUserNameOrPassword.MatchString(testCases[i])
		ast.Equal(expected[i], ret)
	}
}

func TestReDatetime(t *testing.T) {
	ast := assert.New(t)
	testCases := []string{
		// true
		"2019-02-02 10:03:04", "2000-01-01 00:00:00", "2099-12-31 23:59:59",
		// false
		"1999-12-31 23:59:59", "2099-12-31 23:59:60",
	}
	expected := []bool{
		true, true, true,
		false, false,
	}
	for i := range testCases {
		ret := ReDatetime.MatchString(testCases[i])
		ast.Equal(expected[i], ret)
	}
}
