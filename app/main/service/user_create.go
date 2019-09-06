package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/pkg/checker"
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
		return nil, err
	}
	err = dao.MysqlDB.Create(user).Error
	if err != nil {
		return nil, errors.New("用户已存在")
	}
	return user, nil
}
