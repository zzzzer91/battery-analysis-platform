package controller

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func GinResponse(c *gin.Context, s service.Servicer) {
	if res, err := s.Do(); err != nil {
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, res)
	}
}
