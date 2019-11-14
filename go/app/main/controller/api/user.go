package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var s service.UserCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.GinResponse(c, &s)
}

func ModifyUser(c *gin.Context) {
	var s service.UserModifyService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	s.UserName = c.Param("name")
	controller.GinResponse(c, &s)
}

func ListUser(c *gin.Context) {
	var s service.UserListService
	controller.GinResponse(c, &s)
}
