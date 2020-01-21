package dao

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMiningTask(id, name, dataComeFrom, dateRange string) (*model.MiningTask, error) {
	task := model.NewMiningTask(id, name, dataComeFrom, dateRange)
	err := creatTask(consts.MongoCollectionMiningTask, task)
	return task, err
}

func DeleteMiningTask(id string) error {
	return deleteTask(consts.MongoCollectionMiningTask, id)
}

func ListMiningTask() ([]model.MiningTask, error) {
	collection := db.Mongo.Collection(consts.MongoCollectionMiningTask)
	filter := bson.M{}                                // 过滤记录
	projection := bson.M{"_id": false, "data": false} // 过滤字段
	sort := bson.M{"createTime": -1}                  // 结果排序
	// 注意 ctx 不能几个连接复用，原因见 `context.WithTimeout` 源码
	ctx := NewTimeoutCtx()
	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.MiningTask, 0)
	for cur.Next(ctx) {
		result := model.MiningTask{}
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
	collection := db.Mongo.Collection(consts.MongoCollectionMiningTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "data": true} // 注意 _id 默认会返回，需要手动过滤
	ctx := NewTimeoutCtx()
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
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["data"].(bson.A), nil
}
