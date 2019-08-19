package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var code int
	var msg string
	var data interface{}

	if users, err := service.GetUsers(); err != nil {
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
	var userCreateService service.UserCreateService
	if err := c.ShouldBindJSON(&userCreateService); err != nil {
		c.AbortWithError(500, err)
		return
	} else {
		var code int
		var msg string
		var data interface{}
		user, err := userCreateService.CreateUser()
		if err != nil {
			code = jd.ERROR
			msg = err.Error()
		} else {
			code = jd.SUCCESS
			msg = "创建用户 " + user.Name + " 成功"
		}
		res := jd.Build(code, msg, data)
		c.JSON(200, res)
	}
}

func ModifyUser(c *gin.Context) {

}
