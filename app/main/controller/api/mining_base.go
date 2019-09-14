package api

import (
	"battery-analysis-platform/app/main/service"
	"battery-analysis-platform/pkg/jd"
	"github.com/gin-gonic/gin"
)

func GetBasicData(c *gin.Context) {
	var s service.MiningBaseService
	if err := c.ShouldBindQuery(&s); err != nil {
		c.AbortWithError(500, err)
		return
	}

	data, err := s.Query()
	code, msg := jd.HandleError(err)
	if code == jd.SUCCESS {
		msg = "查询成功"
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
