package service

import (
	"battery-analysis-platform/app/main/dao"
	"battery-analysis-platform/pkg/jd"
)

type ChangePasswordService struct {
	UserName string
	Password string `json:"password"`
}

func (s *ChangePasswordService) Do() (*jd.Response, error) {
	err := dao.ChangeUserPassword(s.UserName, s.Password)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "修改密码成功", nil), nil
}