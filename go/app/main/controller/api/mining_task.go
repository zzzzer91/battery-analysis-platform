package api

import (
	"battery-analysis-platform/app/main/controller"
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateMiningTask(c *gin.Context) {
	s := service.CreateMiningTaskService{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	controller.JsonResponse(c, &s)
}

func GetMiningTaskList(c *gin.Context) {
	s := service.GetMiningTaskListService{}
	controller.JsonResponse(c, &s)
}

func GetMiningTaskData(c *gin.Context) {
	s := service.GetMiningTaskDataService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}

func DeleteMiningTask(c *gin.Context) {
	s := service.DeleteMiningTaskService{
		Id: c.Param("taskId"),
	}
	controller.JsonResponse(c, &s)
}
