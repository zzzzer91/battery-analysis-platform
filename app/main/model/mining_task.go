package model

// 不包含任务数据，这个 model 用来做任务列表的元素，所以不需要
type MiningTask struct {
	Id            string `json:"taskId" bson:"_id"`
	TaskName      string `json:"taskName" bson:"taskName"`
	DataComeFrom  string `json:"dataComeFrom" bson:"dataComeFrom"`
	RequestParams string `json:"requestParams" bson:"requestParams"`
	CreateTime    string `json:"createTime" bson:"createTime"`
	TaskStatus    string `json:"taskStatus" bson:"taskStatus"`
	Comment       string `json:"comment" bson:"comment"`
}

var MiningSupportTaskSet map[string]struct{}

func init() {
	MiningSupportTaskSet = map[string]struct{}{
		"充电过程": {},
		"工况":   {},
		"电池统计": {},
	}
}
