package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ShowMiningBaseData(c *gin.Context) {
	var s service.MiningBaseShowDataService
	if err := c.ShouldBindQuery(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.GinResponse(c, &s)
}
