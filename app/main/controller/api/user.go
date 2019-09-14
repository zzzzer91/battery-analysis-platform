package api

import (
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data, err := service.GetUserList()
	code, msg := jd.HandleError(err)
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}

func CreateUser(c *gin.Context) {
	var s service.UserCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	_, err := s.CreateUser()
	code, msg := jd.HandleError(err)
	if code == jd.SUCCESS {
		msg = "创建用户 " + s.UserName + " 成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}

func ModifyUser(c *gin.Context) {
	var s service.UserModifyService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	userName := c.Param("name")
	_, err := s.ModifyUser(userName)
	code, msg := jd.HandleError(err)
	if code == jd.SUCCESS {
		msg = "修改用户 " + userName + " 成功"
	}
	res := jd.Build(code, msg, nil)
	c.JSON(200, res)
}
