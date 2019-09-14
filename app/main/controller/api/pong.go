package api

import (
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, jd.Build(jd.SUCCESS, "", "Pong!"))
}
