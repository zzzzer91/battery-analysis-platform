package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/app/main/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	collectionNameTaskList = "mining_tasks"
)

func GetTaskList() ([]model.MiningTask, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	filter := bson.M{}                  // 过滤记录
	projection := bson.M{"data": false} // 过滤字段
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	var records []model.MiningTask
	for cur.Next(ctx) {
		result := model.MiningTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	return records, nil
}

func GetTask(taskId string) (bson.A, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	collection := dao.MongoDB.Collection(collectionNameTaskList)
	filter := bson.M{"_id": taskId}
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
		return nil, err
	}
	return result["data"].(bson.A), nil
}
