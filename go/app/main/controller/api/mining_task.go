package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateMiningTask(c *gin.Context) {
	var s service.MiningTaskCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.GinResponse(c, &s)
}

func DeleteMiningTask(c *gin.Context) {
	var s service.MiningTaskDeleteService
	s.Id = c.Param("taskId")
	controller.GinResponse(c, &s)
}

func ListMiningTask(c *gin.Context) {
	var s service.MiningTaskListService
	controller.GinResponse(c, &s)
}

func ShowMiningTaskData(c *gin.Context) {
	var s service.MiningTaskShowDataService
	s.Id = c.Param("taskId")
	controller.GinResponse(c, &s)
}
