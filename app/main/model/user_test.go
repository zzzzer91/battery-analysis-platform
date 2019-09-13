package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser(t *testing.T) {
	ast := assert.New(t)
	user := &User{
		Id:     1,
		Name:   "xiaoming",
		Type:   UserTypeSuperUser,
		Status: 0,
	}
	err := user.SetPassword("123456")
	ast.Nil(err)

	ast.Equal(true, user.CheckPassword("123456"))
	ast.Equal(false, user.CheckStatusOk())

	expected := `{"userName":"xiaoming","userType":64,"avatarName":null,"lastLoginTime":null,"comment":"","loginCount":0,"userStatus":0,"createTime":null}`
	ret, err := json.Marshal(user)
	ast.Nil(err)
	ast.Equal(expected, string(ret))
}
