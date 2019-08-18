package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/pkg/checker"
	"errors"
)

type UserCreateService struct {
	UserName string `json:"userName" binding:"required,min=5,max=14"`
	Password string `json:"password" binding:"required,min=6,max=14"`
	Comment  string `json:"comment" binding:"max=64"`
}

func (s *UserCreateService) CreateUser() error {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return errors.New("用户名不符合要求")
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return errors.New("密码不符合要求")
	}
	user := &model.User{Name: s.UserName, Comment: s.Comment}
	err := user.SetPassword(s.Password)
	if err != nil {
		return errors.New("")
	}
	err = dao.MysqlDB.Create(user).Error
	if err != nil {
		return errors.New("用户已存在")
	}
	return nil
}
