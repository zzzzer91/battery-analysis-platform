package service

import (
	"battery-analysis-platform/app/main/dao"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
	"battery-analysis-platform/pkg/jtime"
)

type LoginService struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

func (s *LoginService) Do() (*jd.Response, error) {
	if !checker.ReUserNameOrPassword.MatchString(s.UserName) {
		return jd.Err("用户名不合法"), nil
	}
	if !checker.ReUserNameOrPassword.MatchString(s.Password) {
		return jd.Err("密码不合法"), nil
	}

	user, err := dao.GetUser(s.UserName)
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
	err = dao.SaveUserLoginTimeAndCount(user)
	if err != nil {
		return nil, err
	}
	err = dao.SaveUserToCache(user)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "登录成功", user), nil
}

type LoginByCookieService struct {
	UserName string
}

func (s *LoginByCookieService) Do() (*jd.Response, error) {
	user, err := dao.GetUserFromCache(s.UserName)
	if err != nil {
		return nil, err
	}
	if !user.CheckStatusOk() {
		return jd.Err("该用户已被禁止登录"), nil
	}
	return jd.Build(jd.SUCCESS, "", user), nil
}

type LogoutService struct {
	UserName string
}

func (s *LogoutService) Do() (*jd.Response, error) {
	err := dao.DeleteUserFromCache(s.UserName)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", nil), nil
}
