package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	s := service.CreateUserService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func GetUserList(c *gin.Context) {
	s := service.GetCommonUserListService{}
	controller.JsonResponse(c, &s)
}

func UpdateUserInfo(c *gin.Context) {
	s := service.UpdateUserInfoService{
		UserName: c.Param("name"),
	}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}
