package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
)

type UserCreateService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Comment  string `json:"comment"`
}

func (s *UserCreateService) Do() (*jd.Response, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return jd.Err("用户名不符合要求"), nil
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return jd.Err("密码不符合要求"), nil
	}
	if len(s.Comment) > 64 {
		return jd.Err("参数 comment 不合法"), nil
	}

	data, err := model.CreateUser(s.UserName, s.Password, s.Comment)
	if err != nil {
		return jd.Err("用户 " + s.UserName + " 已存在"), nil
	}
	return jd.Build(jd.SUCCESS, "创建用户 "+s.UserName+" 成功", data), nil
}

type UserModifyService struct {
	UserName string
	Comment  string `json:"comment"`
	Status   int    `json:"userStatus"`
}

func (s *UserModifyService) Do() (*jd.Response, error) {
	if s.Status != model.UserStatusNormal && s.Status != model.UserStatusForbiddenLogin {
		return jd.Err("参数 status 不合法"), nil
	}
	if len(s.Comment) > 64 {
		return jd.Err("参数 comment 不合法"), nil
	}

	user, err := model.GetUser(s.UserName)
	if err != nil {
		return nil, err
	}
	user.Comment = s.Comment
	user.Status = s.Status
	err = model.SaveUserChange(user)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "修改用户 "+s.UserName+" 成功", user), nil
}

type UserListService struct {
}

func (UserListService) Do() (*jd.Response, error) {
	userList, err := model.ListCommonUser()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", userList), nil
}
