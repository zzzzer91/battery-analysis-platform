package service

import (
	"battery-analysis-platform/app/main/model"
	"github.com/shirou/gopsutil/mem"
)

func GetSysInfo() (*model.SysInfo, error) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	s := &model.SysInfo{
		Memory: &model.Memory{
			Total:       vm.Total,
			Free:        vm.Free,
			UsedPercent: vm.UsedPercent,
		},
	}
	return s, nil
}
