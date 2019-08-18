package model

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// 用户类型 Type
const (
	// 普通用户
	TypeCommonUser = 0
	// 超级用户
	TypeSuperUser = 1
)

// 用户状态 Status
const (
	StatusForbiddenLogin = 0
	StatusNormal         = 1
)

type User struct {
	ID            int        `json:"-" gorm:"primary_key"`
	Name          string     `json:"name" gorm:"type:char(16);unique;not null"`
	Password      string     `json:"-" gorm:"type:varchar(128);not null"`
	Type          int        `json:"type" gorm:"type:tinyint;not null;default:0"`
	AvatarName    string     `json:"avatarName" gorm:"type:varchar(256);default:'null.jpg'"`
	LastLoginTime *time.Time `json:"lastLoginTime"`
	Comment       string     `json:"-" gorm:"type:varchar(64)"`
	LoginCount    int        `json:"loginCount" gorm:"type:integer;not null;default:0"`
	Status        int        `json:"-" gorm:"type:tinyint;not null;default:1"`
	CreateTime    *time.Time `json:"-" gorm:"default:current_timestamp"`
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
