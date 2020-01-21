package service

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/dao"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
	"strings"
)

type MiningBaseShowDataService struct {
	DataComeFrom string `form:"dataComeFrom"`
	StartDate    string `form:"startDate"`
	NeedParams   string `form:"needParams"`
	DataLimit    int    `form:"dataLimit"`
}

func (s *MiningBaseShowDataService) Do() (*jd.Response, error) {
	// 校验字段合法性
	table, ok := consts.BatteryNameToTable[s.DataComeFrom]
	if !ok {
		return jd.Err("参数 dataComeFrom 不合法"), nil
	}
	fields := strings.Split(s.NeedParams, ",")
	for _, field := range fields {
		_, ok = table.FieldSet[field]
		if !ok {
			return jd.Err("参数 needParams 不合法"), nil
		}
	}
	if !checker.ReDatetime.MatchString(s.StartDate) {
		return jd.Err("参数 startDate 不合法"), nil
	}
	if s.DataLimit > 10000 {
		return jd.Err("参数 dataLimit 不合法"), nil
	}

	data, err := dao.GetBatteryData(table.Name, s.StartDate, s.DataLimit, fields)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return jd.Err("未查询到相关数据"), nil
	}
	return jd.Build(jd.SUCCESS, "查询成功", data), nil
}
