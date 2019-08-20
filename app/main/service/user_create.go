package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/pkg/checker"
	"errors"
)

type UserCreateService struct {
	UserLoginService
	Comment string `json:"comment" binding:"max=64"`
}

func (s *UserCreateService) CreateUser() (*model.User, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return nil, errors.New("用户名不符合要求")
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return nil, errors.New("密码不符合要求")
	}
	user := &model.User{Name: s.UserName, Comment: s.Comment}
	err := user.SetPassword(s.Password)
	if err != nil {
		return nil, errors.New("")
	}
	err = dao.MysqlDB.Create(user).Error
	if err != nil {
		return nil, errors.New("用户已存在")
	}
	return user, nil
}
