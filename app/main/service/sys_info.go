package service

import (
	"battery-anlysis-platform/app/main/model"
	"github.com/shirou/gopsutil/mem"
)

func GetSysInfo() *model.SysInfo {
	vm, _ := mem.VirtualMemory()
	return &model.SysInfo{
		Memory: &model.Memory{
			Total:       vm.Total,
			Free:        vm.Free,
			UsedPercent: vm.UsedPercent,
		},
	}
}
