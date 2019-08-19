package model

import "github.com/shirou/gopsutil/mem"

type Memory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type SysInfo struct {
	Memory *Memory `json:"memory"`
}

func NewSysInfo() *SysInfo {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return nil
	}
	return &SysInfo{
		Memory: &Memory{
			Total:       vm.Total,
			Free:        vm.Free,
			UsedPercent: vm.UsedPercent,
		},
	}
}
