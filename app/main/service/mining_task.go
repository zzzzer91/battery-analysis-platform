package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"battery-anlysis-platform/app/main/worker"
	"battery-anlysis-platform/pkg/checker"
	"battery-anlysis-platform/pkg/jtime"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	collectionNameTaskList = "mining_tasks"
	timeout                = time.Second

	// support task
	taskBasename         = "task."
	taskComputeModel     = taskBasename + "compute_model"
	taskStopComputeModel = taskBasename + "stop_compute_model"
)

type MiningCreateTaskService struct {
	TaskName     string `json:"taskName" binding:"required"`
	DataComeFrom string `json:"dataComeFrom" binding:"required"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	AllData      bool   `json:"allData"` // bool 型不能 required，因为 false 会被判空
}

func (s *MiningCreateTaskService) CreateTask() (*model.MiningTask, error) {
	// needFields 的顺序不能变，追加新字段，必须放在最后
	var needFields string
	switch s.TaskName {
	case "充电过程":
		needFields = "bty_t_vol, bty_t_curr, battery_soc, id, byt_ma_sys_state"
	case "工况":
		needFields = "timestamp, bty_t_curr, met_spd"
	case "电池统计":
		needFields = "max_t_s_b_num, min_t_s_b_num"
	default:
		return nil, errors.New("参数 TaskName 不合法")
	}
	table, ok := model.BatteryMysqlNameToTable[s.DataComeFrom]
	if !ok {
		return nil, errors.New("参数 dataComeFrom 不合法")
	}
	var requestParams string
	if s.AllData {
		requestParams = "所有数据"
	} else {
		if !checker.ReDatetime.MatchString(s.StartDate) {
			return nil, errors.New("参数 startDate 不合法")
		}
		if !checker.ReDatetime.MatchString(s.EndDate) {
			return nil, errors.New("参数 EndDate 不合法")
		}
		requestParams = s.StartDate + " - " + s.EndDate
	}

	asyncResult, err := worker.Celery.Delay(
		taskComputeModel,
		s.TaskName, needFields, table.Name, requestParams)
	if err != nil {
		panic(err)
	}

	task := &model.MiningTask{
		Id:            asyncResult.TaskID,
		TaskName:      s.TaskName,
		DataComeFrom:  s.DataComeFrom,
		RequestParams: requestParams,
		CreateTime:    jtime.NowStr(),
		TaskStatus:    "执行中",
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	_, err = collection.InsertOne(ctx, &task)
	if err != nil {
		// 终止正在执行的任务
		_, _ = worker.Celery.Delay(taskStopComputeModel, task.Id)
		panic(err)
	}

	return task, nil
}

func GetTaskList() ([]model.MiningTask, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	filter := bson.M{}                  // 过滤记录
	projection := bson.M{"data": false} // 过滤字段
	sort := bson.M{"createTime": -1}    // 结果排序
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		panic(err)
	}
	var records []model.MiningTask
	for cur.Next(ctx) {
		result := model.MiningTask{}
		err := cur.Decode(&result)
		if err != nil {
			panic(err)
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetTask(id string) (bson.A, error) {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	filter := bson.M{"_id": id}
	projection := bson.M{"_id": false, "data": true} // 注意 _id 默认会返回，需要手动过滤
	// 注意 bson.E 不能用来映射 mongo 中的 map，
	// 要么使用 bson.D，采用 []bson.E 代表一个字典，其中 bson.E 是 struct，有 key 和 value 字段，
	// 此时，映射出来的子字典也都是 bson.D 类型，
	// 而映射出来的列表是 bson.A 类型，
	// bson.D 在 JSON 序列化时会在最外层加上 []，所以需要序列化的结果不要用，而采用 bson.M；
	// 要么使用 bson.M，采用 map[string]interface{} 代表一个字典，
	// 此时，映射出来的子字典也都是 bson.M 类型，
	// 而映射出来的列表也是 bson.A 类型，
	// 这种方法 JSON 序列化时符合直觉，推荐使用；
	// 若要代表一个列表，类似 Python 中 list，不限定类型，使用 bson.A，即 []interface{}。
	var result bson.M
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result["data"].(bson.A), nil
}

func DeleteTask(id string) (int64, error) {
	// 因为 gocelery 未提供终止任务的 api，这里把终止行为封装成任务，然后调用它
	_, err := worker.Celery.Delay(taskStopComputeModel, id)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	filter := bson.M{"_id": id}
	ret, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	return ret.DeletedCount, nil
}
