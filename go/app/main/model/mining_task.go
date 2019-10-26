package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 过滤不合法任务
var MiningSupportTaskSet = map[string]struct{}{
	"充电过程":        {},
	"工况":          {},
	"电池统计":        {},
	"pearson相关系数": {},
}

// 不包含任务数据，这个 model 用来做任务列表的元素，所以不需要
type MiningTask struct {
	TaskId       string `json:"taskId" bson:"taskId"`
	TaskName     string `json:"taskName" bson:"taskName"`
	DataComeFrom string `json:"dataComeFrom" bson:"dataComeFrom"`
	DateRange    string `json:"dateRange" bson:"dateRange"`
	CreateTime   string `json:"createTime" bson:"createTime"`
	TaskStatus   string `json:"taskStatus" bson:"taskStatus"`
	Comment      string `json:"comment" bson:"comment"`
}

func CreateMiningTask(id, name, dataComeFrom, dateRange string) (*MiningTask, error) {
	collection := db.Mongo.Collection(mongoCollectionMiningTasks)
	task := MiningTask{
		TaskId:       id,
		TaskName:     name,
		DataComeFrom: dataComeFrom,
		DateRange:    dateRange,
		CreateTime:   jtime.NowStr(),
		TaskStatus:   "执行中",
	}
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func ListMiningTask() ([]MiningTask, error) {
	collection := db.Mongo.Collection(mongoCollectionMiningTasks)
	filter := bson.M{}                  // 过滤记录
	projection := bson.M{"data": false} // 过滤字段
	sort := bson.M{"createTime": -1}    // 结果排序
	// 注意 ctx 不能几个连接复用
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]MiningTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), mongoCtxTimeout)
	for cur.Next(ctx) {
		result := MiningTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetMiningTaskData(id string) (bson.A, error) {
	collection := db.Mongo.Collection(mongoCollectionMiningTasks)
	filter := bson.M{"taskId": id}
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
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["data"].(bson.A), nil
}

func DeleteMiningTask(id string) (int64, error) {
	collection := db.Mongo.Collection(mongoCollectionMiningTasks)
	filter := bson.M{"taskId": id}
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	ret, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}
