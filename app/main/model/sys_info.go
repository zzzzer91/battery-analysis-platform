package model

type Memory struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}

type SysInfo struct {
	Memory *Memory `json:"memory"`
}
