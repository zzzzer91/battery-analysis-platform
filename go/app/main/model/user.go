package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"battery-analysis-platform/pkg/security"
	"encoding/json"
	"time"
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

const (
	// redis 的键前缀
	redisKeyPrefix = "ubattery:users:"
	// 缓存中 key 过期时间
	redisKeyExpiration = time.Hour
)

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"userName"`
	Password string `json:"-"`
	Type     int    `json:"userType"`
	// *string 让其 json 时可以返回 null，否则只能返回字符串零值
	AvatarName    *string     `json:"avatarName"`
	LastLoginTime *jtime.Time `json:"lastLoginTime"`
	Comment       string      `json:"comment"`
	LoginCount    int         `json:"loginCount"`
	Status        int         `json:"userStatus"`
	CreateTime    *jtime.Time `json:"createTime"`
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
	user := User{Name: name, Comment: comment}
	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}
	// 这里必须传入指针
	err = db.Gorm.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
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

// ---------------------------cache---------------------------

func SaveUserToCache(user *User) error {
	// 存储 JSON 序列化的数据
	jd, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return db.Redis.Set(redisKeyPrefix+user.Name, jd, redisKeyExpiration).Err()
}

func DeleteUserFromCache(name string) error {
	return db.Redis.Del(redisKeyPrefix + name).Err()
}

func GetUserFromCache(name string) (*User, error) {
	val, err := db.Redis.Get(redisKeyPrefix + name).Bytes()
	if err != nil {
		return nil, err
	}
	user := User{}
	err = json.Unmarshal(val, &user)
	if err != nil {
		return nil, err
	}
	// 刷新 key 的过期时间
	db.Redis.Expire(redisKeyPrefix+name, redisKeyExpiration)
	return &user, nil
}
