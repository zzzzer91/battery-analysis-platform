package service

import (
	"battery-anlysis-platform/app/main/model"
	"errors"
)

func GetSysInfo() (*model.SysInfo, error) {
	ret := model.NewSysInfo()
	if ret == nil {
		return nil, errors.New("")
	}
	return ret, nil
}
