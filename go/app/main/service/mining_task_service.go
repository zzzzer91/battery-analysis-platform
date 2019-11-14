package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/producer"
	"battery-analysis-platform/pkg/checker"
	"battery-analysis-platform/pkg/jd"
)

// support task
const (
	miningTaskComputeModel     = "task.mining.compute_model"
	miningTaskStopComputeModel = "task.mining.stop_compute_model"
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

	asyncResult, err := producer.Celery.Delay(
		miningTaskComputeModel,
		s.TaskName, table.Name, dateRange)
	if err != nil {
		return nil, err
	}

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
	_, err := producer.Celery.Delay(miningTaskStopComputeModel, s.Id)
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
