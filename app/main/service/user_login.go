package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"errors"
	"time"
)

type UserLoginService struct {
	UserName string `json:"userName" binding:"required,min=5,max=14"`
	Password string `json:"password" binding:"required,min=6,max=14"`
}

func (s *UserLoginService) Login() (*model.User, error) {
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
	t := time.Now()
	user.LastLoginTime = &t
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
