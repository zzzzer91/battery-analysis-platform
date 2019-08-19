package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"errors"
)

func GetUsers() ([]model.User, error) {
	var users []model.User
	err := dao.MysqlDB.Where("type != ?", model.TypeSuperUser).Find(&users).Error
	if err != nil {
		return nil, errors.New("")
	}
	return users, nil
}
