package api

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func ShowSysInfo(c *gin.Context) {
	var s service.SysInfoShowService
	res, err := s.Do()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, res)
}
