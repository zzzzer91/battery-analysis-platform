package api

import (
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetBasicData(c *gin.Context) {
	var s service.MiningBaseService
	if err := c.ShouldBindQuery(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	var code int
	var msg string
	data, err := s.Query()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "查询成功"
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
