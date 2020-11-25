package service

import (
	"battery-analysis-platform/app/web/cache"
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/dal"
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/app/web/producer"
	"battery-analysis-platform/pkg/conv"
	"battery-analysis-platform/pkg/jd"
	"fmt"
)

type CreateDlTaskService struct {
	Dataset        string                  `json:"dataset"`
	HyperParameter *model.NnHyperParameter `json:"hyperParameter"`
}

func (s *CreateDlTaskService) Do() (*jd.Response, error) {
	// TODO 检查输入参数

	// 检查是否达到创建任务上限
	if !cache.GetRedisService().CheckTaskLimit(constant.RedisKeyDlTaskWorkingIdSet, 1) {
		return jd.Err("允许同时执行任务数已达上限"), nil
	}

	taskID, err := producer.GetCeleryService().Delay(
		constant.CeleryTaskDeeplearningTrain, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}
	// 添加正在工作的任务的 id 到集合中
	err = cache.GetRedisService().AddWorkingTaskIdToSet(constant.RedisKeyDlTaskWorkingIdSet, taskID)
	if err != nil {
		return nil, err
	}

	data, err := dal.GetMongoService().CreateDlTask(taskID, s.Dataset, s.HyperParameter)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "创建成功", data), nil
}

type GetDlTaskListService struct {
}

func (s *GetDlTaskListService) Do() (*jd.Response, error) {
	data, err := dal.GetMongoService().GetDlTaskList()
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type GetDlTaskTraningHistoryService struct {
	Id            string
	ReadFromRedis bool
}

func (s *GetDlTaskTraningHistoryService) Do() (*jd.Response, error) {
	var data *model.NnTrainingHistory
	if s.ReadFromRedis {
		prefixStr := constant.RedisPrefixDlTaskTrainingHistory + s.Id + ":"

		lossStrList, err := cache.GetRedisService().LRange(
			prefixStr+"loss", 0, -1)
		if err != nil {
			return nil, err
		}
		// 转换为 float
		lossList, err := conv.StringSlice2FloatSlice(lossStrList)
		if err != nil {
			return nil, err
		}

		accuracyStrList, err := cache.GetRedisService().LRange(
			prefixStr+"accuracy", 0, -1)
		if err != nil {
			return nil, err
		}
		accuracyList, err := conv.StringSlice2FloatSlice(accuracyStrList)
		if err != nil {
			return nil, err
		}

		data = &model.NnTrainingHistory{
			Loss:     lossList,
			Accuracy: accuracyList,
		}
	} else {
		var err error
		data, err = dal.GetMongoService().GetDlTaskTrainingHistory(s.Id)
		if err != nil {
			return nil, err
		}
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type GetDlTaskEvalResultService struct {
	Id string
}

func (s *GetDlTaskEvalResultService) Do() (*jd.Response, error) {
	data, err := dal.GetMongoService().GetDlTaskEvalResult(s.Id)
	if err != nil {
		return nil, err
	}
	return jd.Build(jd.SUCCESS, "", data), nil
}

type DownloadDlModelService struct {
	Id string
}

func (s *DownloadDlModelService) Do() (string, error) {
	return conf.App.Gin.ResourcePath + constant.FileDlModelPath + fmt.Sprintf("/%s.pt", s.Id), nil
}

type DeleteDlTaskService struct {
	Id string
}

func (s *DeleteDlTaskService) Do() (*jd.Response, error) {
	// 因为 gocelery 未提供终止任务的 api，这里把终止行为封装成任务，然后调用它
	_, err := producer.GetCeleryService().Delay(constant.CeleryTaskDeeplearningStopTrain, s.Id)
	if err != nil {
		return nil, err
	}

	err = cache.GetRedisService().DelWorkingTaskIdFromSet(constant.RedisKeyDlTaskWorkingIdSet, s.Id)
	if err != nil {
		return nil, err
	}

	// 删除暂存在 redis 中的数据
	prefixStr := constant.RedisPrefixDlTaskTrainingHistory + s.Id + ":"
	cache.GetRedisService().Del(prefixStr+constant.RedisCommonKeySigList, prefixStr+"loss", prefixStr+"accuracy")

	err = dal.GetMongoService().DeleteDlTask(s.Id)
	if err != nil {
		return nil, err
	}

	return jd.Build(jd.SUCCESS, "删除成功", nil), nil
}
