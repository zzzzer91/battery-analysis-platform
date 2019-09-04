package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"errors"
)

type UserCreateService struct {
	UserName string `json:"userName" binding:"required,min=5,max=14"`
	Password string `json:"password" binding:"required,min=5,max=14"`
	Comment  string `json:"comment" binding:"max=64"`
}

func (s *UserCreateService) CreateUser() (*model.User, error) {
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
