package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
	"battery-analysis-platform/pkg/jtime"
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

func (s *UserListService) Do() (*jd.Response, error) {
	userList, err := model.ListCommonUser()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", userList), nil
}

type UserLoginService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *UserLoginService) Do() (*jd.Response, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return jd.Err("用户名不合法"), nil
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return jd.Err("密码不合法"), nil
	}

	user, err := model.GetUser(s.UserName)
	if err != nil {
		return jd.Err("账号或密码错误"), nil
	}
	if !user.CheckPassword(s.Password) {
		return jd.Err("账号或密码错误"), nil
	}
	if !user.CheckStatusOk() {
		return jd.Err("该用户已被禁止登录"), nil
	}
	user.LastLoginTime = jtime.Now()
	user.LoginCount += 1
	err = model.SaveUserChange(user)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "登录成功", user), nil
}

type UserLoginByCookieService struct {
	UserName string
}

func (s *UserLoginByCookieService) Do() (*jd.Response, error) {
	user, err := model.GetUser(s.UserName)
	if err != nil {
		return nil, err
	}
	if !user.CheckStatusOk() {
		return jd.Err("该用户已被禁止登录"), nil
	}
	return jd.Build(jd.SUCCESS, "", user), nil
}

type UserLogoutService struct {
}

func (s *UserLogoutService) Do() (*jd.Response, error) {
	return jd.Build(jd.SUCCESS, "", nil), nil
}
