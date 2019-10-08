package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"battery-analysis-platform/pkg/security"
)

// 用户类型 Type
const (
	// 超级用户
	UserTypeSuperUser = 64
	// 普通用户
	UserTypeCommonUser = 0
)

// 用户状态 Status
const (
	UserStatusForbiddenLogin = 0
	UserStatusNormal         = 1
)

type User struct {
	Id       int    `json:"-" gorm:"primary_key"`
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
	s, err := security.GeneratePasswordHash(password, "pbkdf2:sha256", 8)
	if err != nil {
		return err
	}
	user.Password = s
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := security.CheckPasswordHash(user.Password, password)
	return err == nil
}

func (user *User) CheckStatusOk() bool {
	return user.Status == UserStatusNormal
}

func CreateUser(name, password, comment string) (*User, error) {
	user := &User{Name: name, Comment: comment}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	err = db.Gorm.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(name string) (*User, error) {
	var user User
	// 找不到会返回 ErrRecordNotFound 错误
	err := db.Gorm.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func ListCommonUser() ([]User, error) {
	var users []User
	// 找不到会返回空列表，不会返回错误
	err := db.Gorm.Where("type != ?", UserTypeSuperUser).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func SaveUserChange(user *User) error {
	return db.Gorm.Save(user).Error
}
