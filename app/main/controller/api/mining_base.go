package api

import (
	"battery-anlysis-platform/app/main/mapping"
	"battery-anlysis-platform/app/main/service"
	"battery-anlysis-platform/pkg/checker"
	"battery-anlysis-platform/pkg/jd"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GetBasicData(c *gin.Context) {
	dataComeFrom := c.Query("dataComeFrom")
	if dataComeFrom == "" {
		c.AbortWithError(500, errors.New("URL 参数 dataComeFrom 为空"))
		return
	}
	needParams := c.Query("needParams")
	if needParams == "" {
		c.AbortWithError(500, errors.New("URL 参数 needParams 为空"))
		return
	}
	startDate := c.Query("startDate")
	if startDate == "" {
		c.AbortWithError(500, errors.New("URL 参数 startDate 为空"))
		return
	}
	dataLimitStr := c.Query("dataLimit")
	if dataLimitStr == "" {
		c.AbortWithError(500, errors.New("URL 参数 dataLimit 为空"))
		return
	}

	// 校验字段合法性
	table, ok := mapping.MysqlNameToTable[dataComeFrom]
	if !ok {
		c.AbortWithError(500, errors.New("dataComeFrom 表不存在"))
		return
	}
	fields := strings.Split(needParams, ",")
	for _, field := range fields {
		_, ok = table.FieldToName[field]
		if !ok {
			c.AbortWithError(500, errors.New("needParams 中有字段不存在"))
			return
		}
	}
	if !checker.ReDatetime.MatchString(startDate) {
		c.AbortWithError(500, errors.New("startDate 不合法"))
		return
	}
	dataLimit, err := strconv.Atoi(dataLimitStr)
	if err != nil || dataLimit > 10000 {
		c.AbortWithError(500, errors.New("dataLimit 不合法"))
		return
	}

	s := service.MiningBaseService{
		TableName: table.Name,
		StartDate: startDate,
		Fields:    fields,
		DataLimit: dataLimit,
	}

	var code int
	var msg string
	var data interface{}

	ret, err := s.Query()
	if err != nil {
		code = jd.ERROR
		msg = err.Error()
	} else {
		code = jd.SUCCESS
		msg = "查询成功"
		data = ret
	}
	res := jd.Build(code, msg, data)
	c.JSON(200, res)
}
