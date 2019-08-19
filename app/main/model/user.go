package model

import (
	"battery-anlysis-platform/pkg/jtime"
	"golang.org/x/crypto/bcrypt"
)

// 用户类型 Type
const (
	// 超级用户
	TypeSuperUser = 64
	// 普通用户
	TypeCommonUser = 0
)

// 用户状态 Status
const (
	StatusForbiddenLogin = 0
	StatusNormal         = 1
)

type User struct {
	ID       int    `json:"-" gorm:"primary_key"`
	Name     string `json:"userName" gorm:"type:char(16);unique_index;not null"`
	Password string `json:"-" gorm:"type:varchar(128);not null"`
	Type     int    `json:"userType" gorm:"type:tinyint;not null;default:0"`
	// *string 让其 json 时可以返回 null，否则只能返回字符串零值
	AvatarName    *string         `json:"avatarName" gorm:"type:varchar(256)"`
	LastLoginTime *jtime.JSONTime `json:"lastLoginTime" gorm:"type:datetime"`
	Comment       string          `json:"comment" gorm:"type:varchar(64)"`
	LoginCount    int             `json:"loginCount" gorm:"type:integer;not null;default:0"`
	Status        int             `json:"userStatus" gorm:"type:tinyint;not null;default:1"`
	CreateTime    *jtime.JSONTime `json:"createTime" gorm:"type:datetime;default:current_timestamp"`
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) CheckStatusOk() bool {
	return user.Status == StatusNormal
}
