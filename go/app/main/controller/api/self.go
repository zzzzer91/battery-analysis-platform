package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func UpdateSelfPassword(c *gin.Context) {
	val, _ := c.Get("user")
	user := val.(*model.User)
	s := service.UpdateUserPasswordService{
		UserName: user.Name,
	}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}
