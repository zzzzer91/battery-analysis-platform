package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"errors"
)

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
