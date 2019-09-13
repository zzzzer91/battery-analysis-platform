package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/pkg/checker"
	"battery-anlysis-platform/pkg/jtime"
	"errors"
)

type UserCreateService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Comment  string `json:"comment"`
}

func (s *UserCreateService) CreateUser() (*model.User, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return nil, errors.New("用户名不符合要求")
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return nil, errors.New("密码不符合要求")
	}
	if len(s.Comment) > 64 {
		return nil, errors.New("参数 comment 不合法")
	}

	user := &model.User{Name: s.UserName, Comment: s.Comment}
	err := user.SetPassword(s.Password)
	if err != nil {
		panic(err)
	}
	err = dao.MysqlDB.Create(user).Error
	if err != nil {
		return nil, errors.New("用户已存在")
	}
	return user, nil
}

type UserModifyService struct {
	Status  int    `json:"userStatus"`
	Comment string `json:"comment"`
}

func (s *UserModifyService) ModifyUser(name string) (*model.User, error) {
	if s.Status != model.UserStatusNormal && s.Status != model.UserStatusForbiddenLogin {
		return nil, errors.New("参数 status 不合法")
	}
	if len(s.Comment) > 64 {
		return nil, errors.New("参数 comment 不合法")
	}

	var user model.User
	err := dao.MysqlDB.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	user.Comment = s.Comment
	user.Status = s.Status
	dao.MysqlDB.Save(&user)
	return &user, nil
}

type UserLoginService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *UserLoginService) Login() (*model.User, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return nil, errors.New("用户名不合法")
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return nil, errors.New("密码不合法")
	}

	var user model.User
	err := dao.MysqlDB.Where("name = ?", s.UserName).First(&user).Error
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}
	if !user.CheckPassword(s.Password) {
		return nil, errors.New("账号或密码错误")
	}
	if !user.CheckStatusOk() {
		return nil, errors.New("该用户已被禁止登录")
	}
	user.LastLoginTime = jtime.Now()
	user.LoginCount += 1
	dao.MysqlDB.Save(&user)
	return &user, nil
}

func LoginByCookie(id int) (*model.User, error) {
	var user model.User
	err := dao.MysqlDB.First(&user, id).Error
	if err != nil {
		return nil, errors.New("")
	}
	if !user.CheckStatusOk() {
		return nil, errors.New("该用户已被禁止登录")
	}
	return &user, nil
}

func GetUserList() ([]model.User, error) {
	var users []model.User
	err := dao.MysqlDB.Where("type != ?", model.UserTypeSuperUser).Find(&users).Error
	if err != nil {
		return nil, errors.New("")
	}
	return users, nil
}
