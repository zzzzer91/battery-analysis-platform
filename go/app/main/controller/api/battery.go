package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func GetBatteryList(c *gin.Context) {
	s := service.GetBatteryListService{}
	if err := c.ShouldBindQuery(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}
