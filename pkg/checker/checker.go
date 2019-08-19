package checker

import (
	"regexp"
)

// 过滤不合法用户名和密码
var ReUserNameOrPassword = regexp.MustCompile(`^[0-9a-zA-Z]{6,14}$`)

// 过滤不合法日期，只能 2000-2099 年
// 润年，2 月 没检查
var ReDatetime = regexp.MustCompile("^20[0-9]{2}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[1-2][0-9]|3[0-1]) (?:[0-1][0-9]|2[0-4]):[0-5][0-9]:[0-5][0-9]$")
