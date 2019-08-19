package checker

import (
	"regexp"
)

var ReUserNameOrPassword *regexp.Regexp
var ReDatetime *regexp.Regexp

func init() {
	pattern := regexp.MustCompile(`^[0-9a-zA-Z]{6,14}$`)
	ReUserNameOrPassword = pattern

	pattern = regexp.MustCompile("^20[0-9]{2}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[1-2][0-9]|3[0-1]) (?:[0-1][0-9]|2[0-4]):[0-5][0-9]:[0-5][0-9]$")
	ReDatetime = pattern
}
