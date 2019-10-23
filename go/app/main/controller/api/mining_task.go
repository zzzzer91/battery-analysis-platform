package api

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func CreateMiningTask(c *gin.Context) {
	var s service.MiningTaskCreateService
	if err := c.ShouldBindJSON(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func DeleteMiningTask(c *gin.Context) {
	var s service.MiningTaskDeleteService
	s.Id = c.Param("taskId")
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ListMiningTask(c *gin.Context) {
	var s service.MiningTaskListService
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}

func ShowMiningTaskData(c *gin.Context) {
	var s service.MiningTaskShowDataService
	s.Id = c.Param("taskId")
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}
