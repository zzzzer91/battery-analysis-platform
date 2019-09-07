package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"errors"
)

func GetUserList() ([]model.User, error) {
	var users []model.User
	err := dao.MysqlDB.Where("type != ?", model.UserTypeSuperUser).Find(&users).Error
	if err != nil {
		return nil, errors.New("")
	}
	return users, nil
}
