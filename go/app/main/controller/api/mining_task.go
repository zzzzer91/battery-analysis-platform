package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateMiningTask(c *gin.Context) {
	s := service.MiningTaskCreateService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func DeleteMiningTask(c *gin.Context) {
	s := service.MiningTaskDeleteService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}

func ListMiningTask(c *gin.Context) {
	s := service.MiningTaskListService{}
	controller.JsonResponse(c, &s)
}

func ShowMiningTaskData(c *gin.Context) {
	s := service.MiningTaskShowDataService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}
