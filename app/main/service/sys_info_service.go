package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/pkg/jd"
)

type SysInfoShowService struct {
}

func (SysInfoShowService) Do() (*jd.Response, error) {
	data, err := model.GetSysInfo()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}
