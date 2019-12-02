package controller

import (
	"battery-analysis-platform/app/main/service"
	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, s service.JsonServicer) {
	if res, err := s.Do(); err != nil {
		// 返回 500 错误码，并且会在控制台打印错误
		c.AbortWithError(500, err)
	} else {
		c.JSON(200, res)
	}
}
