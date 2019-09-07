package service

import (
	"testing"
)

func TestGetUsers(t *testing.T) {
	ret, _ := GetUserList()
	t.Log(ret)
}
