package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var code int
	var msg string
	var data interface{}

	if users, err := service.GetUserList(); err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		data = users
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func CreateUser(c *gin.Context) {
	var s service.UserCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	var code int
	var msg string
	if user, err := s.CreateUser(); err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "创建用户 " + user.Name + " 成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}

func ModifyUser(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.AbortWithError(500, errors.New("URL 参数 name 为空"))
		return
	}
	var s service.UserModifyService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	var code int
	var msg string
	if user, err := s.ModifyUser(name); err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "修改用户 " + user.Name + " 成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}
