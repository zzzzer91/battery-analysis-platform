package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/producer"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
)

type MiningTaskCreateService struct {
	TaskName     string `json:"taskName"`
	DataComeFrom string `json:"dataComeFrom"`
	// BatteryStatus int    `json:"batteryStatus"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	AllData   bool   `json:"allData"` // bool 型不能 required，因为 false 会被判空
}

func (s *MiningTaskCreateService) Do() (*jd.Response, error) {
	if _, ok := model.MiningSupportTaskSet[s.TaskName]; !ok {
		return jd.Err("参数 TaskName 不合法"), nil
	}

	table, ok := model.BatteryNameToTable[s.DataComeFrom]
	if !ok {
		return jd.Err("参数 dataComeFrom 不合法"), nil
	}

	var dateRange string
	if s.AllData {
		dateRange = "所有数据"
	} else {
		if !checker.ReDatetime.MatchString(s.StartDate) {
			return jd.Err("参数 startDate 不合法"), nil
		}
		if !checker.ReDatetime.MatchString(s.EndDate) {
			return jd.Err("参数 EndDate 不合法"), nil
		}
		dateRange = s.StartDate + " - " + s.EndDate
	}

	// 检查是否达到创建任务上限
	if !producer.CheckTaskLimit("miningTask:workingIdSet", 1) {
		return jd.Err("允许同时执行任务数已达上限"), nil
	}

	// 调用 celery
	asyncResult, err := producer.Celery.Delay(
		"task.mining.compute_model",
		s.TaskName, table.Name, dateRange)
	if err != nil {
		return nil, err
	}
	// 添加正在工作的任务的 id 到 redis 集合中
	err = producer.AddWorkingTaskIdToSet("miningTask:workingIdSet", asyncResult.TaskID)
	if err != nil {
		return nil, err
	}

	// 创建 mongo 记录
	data, err := model.CreateMiningTask(asyncResult.TaskID, s.TaskName, s.DataComeFrom, dateRange)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "创建成功", data), nil
}

type MiningTaskDeleteService struct {
	Id string
}

func (s *MiningTaskDeleteService) Do() (*jd.Response, error) {
	// 因为 gocelery 未提供终止任务的 api，这里把终止行为封装成任务，然后调用它
	_, err := producer.Celery.Delay("task.mining.stop_compute_model", s.Id)
	if err != nil {
		return nil, err
	}

	err = producer.DelWorkingTaskIdFromSet("miningTask:workingIdSet", s.Id)
	if err != nil {
		return nil, err
	}

	err = model.DeleteMiningTask(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "删除成功", nil), nil
}

type MiningTaskListService struct {
}

func (s *MiningTaskListService) Do() (*jd.Response, error) {
	data, err := model.ListMiningTask()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type MiningTaskShowDataService struct {
	Id string
}

func (s *MiningTaskShowDataService) Do() (*jd.Response, error) {
	data, err := model.GetMiningTaskData(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}
