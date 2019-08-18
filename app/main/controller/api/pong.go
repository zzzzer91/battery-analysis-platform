package api

import (
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, jd.Build(jd.SUCCESS, "", "Pong!"))
}
