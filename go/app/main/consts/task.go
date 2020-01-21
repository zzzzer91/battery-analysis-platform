package consts

import "time"

// 过滤不合法任务
var MiningSupportTaskSet = map[string]struct{}{
	"充电过程":        {},
	"工况":          {},
	"电池统计":        {},
	"pearson相关系数": {},
}

const (
	TaskWaitSigTimeout = time.Second * 3
)
