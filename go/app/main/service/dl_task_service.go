package service

import (
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/producer"
	"battery-analysis-platform/pkg/jd"
)

const (
	dlTaskTrain     = "task.deeplearning.train"
	dlTaskStopTrain = "task.deeplearning.stop_train"
)

type DlTaskCreateService struct {
	Dataset        string                  `json:"dataset"`
	HyperParameter *model.NnHyperParameter `json:"hyperParameter"`
}

func (s *DlTaskCreateService) Do() (*jd.Response, error) {
	// TODO 检查输入参数

	asyncResult, err := producer.Celery.Delay(
		dlTaskTrain, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}

	data, err := model.CreateDlTask(asyncResult.TaskID, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "创建成功", data), nil
}

type DlTaskDeleteService struct {
	Id string
}

func (s *DlTaskDeleteService) Do() (*jd.Response, error) {
	// 因为 gocelery 未提供终止任务的 api，这里把终止行为封装成任务，然后调用它
	_, err := producer.Celery.Delay(dlTaskStopTrain, s.Id)
	if err != nil {
		return nil, err
	}

	_, err = model.DeleteDlTask(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "删除成功", nil), nil
}

type DlTaskListService struct {
}

func (DlTaskListService) Do() (*jd.Response, error) {
	data, err := model.ListDlTask()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}
