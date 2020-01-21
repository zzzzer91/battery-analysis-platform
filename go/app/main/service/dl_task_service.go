package service

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/dao"
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/app/main/producer"
	"battery-analysis-platform/pkg/jd"
	"fmt"
)

type DlTaskCreateService struct {
	Dataset        string                  `json:"dataset"`
	HyperParameter *model.NnHyperParameter `json:"hyperParameter"`
}

func (s *DlTaskCreateService) Do() (*jd.Response, error) {
	// TODO 检查输入参数

	// 检查是否达到创建任务上限
	if !producer.CheckTaskLimit("deeplearningTask:workingIdSet", 1) {
		return jd.Err("允许同时执行任务数已达上限"), nil
	}

	asyncResult, err := producer.Celery.Delay(
		"task.deeplearning.train", s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}
	// 添加正在工作的任务的 id 到集合中
	err = producer.AddWorkingTaskIdToSet("deeplearningTask:workingIdSet", asyncResult.TaskID)
	if err != nil {
		return nil, err
	}

	data, err := dao.CreateDlTask(asyncResult.TaskID, s.Dataset, s.HyperParameter)
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
	_, err := producer.Celery.Delay("task.deeplearning.stop_train", s.Id)
	if err != nil {
		return nil, err
	}

	err = producer.DelWorkingTaskIdFromSet("deeplearningTask:workingIdSet", s.Id)
	if err != nil {
		return nil, err
	}

	// 删除暂存在 redis 中的数据
	prefixStr := "deeplearningTask:trainingHistory:" + s.Id + ":"
	db.Redis.Del(prefixStr+"sigList", prefixStr+"loss", prefixStr+"accuracy")

	err = dao.DeleteDlTask(s.Id)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "删除成功", nil), nil
}

type DlTaskListService struct {
}

func (s *DlTaskListService) Do() (*jd.Response, error) {
	data, err := dao.ListDlTask()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type DlTaskShowTraningHistoryService struct {
	Id            string
	ReadFromRedis bool
}

func (s *DlTaskShowTraningHistoryService) Do() (*jd.Response, error) {
	data, err := dao.GetDlTaskTrainingHistory(s.Id, s.ReadFromRedis)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type DlTaskShowEvalResultService struct {
	Id string
}

func (s *DlTaskShowEvalResultService) Do() (*jd.Response, error) {
	data, err := dao.GetDlTaskEvalResult(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type DlDownloadModelService struct {
	Id string
}

func (s *DlDownloadModelService) Do() (string, error) {
	return conf.App.Gin.ResourcePath + consts.FileDlModelPath + fmt.Sprintf("/%s.pt", s.Id), nil
}
