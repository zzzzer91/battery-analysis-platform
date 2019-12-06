package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	s := service.UserCreateService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func ModifyUser(c *gin.Context) {
	s := service.UserModifyService{
		UserName: c.Param("name"),
	}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func ListUser(c *gin.Context) {
	s := service.UserListService{}
	controller.JsonResponse(c, &s)
}
