package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NnLayer struct {
	Neurons    int    `json:"neurons" bson:"neurons"`
	Activation string `json:"activation" bson:"activation"`
}

type DlTask struct {
	TaskId         string    `json:"taskId" bson:"taskId"`
	Dataset        string    `json:"dataset" bson:"dataset"`
	Loss           string    `json:"loss" bson:"loss"`
	Epochs         int       `json:"epochs" bson:"epochs"`
	BatchSize      int       `json:"batchSize" bson:"batchSize"`
	NnArchitecture []NnLayer `json:"nnArchitecture" bson:"nnArchitecture"`
	CreateTime     string    `json:"createTime" bson:"createTime"`
	TaskStatus     string    `json:"taskStatus" bson:"taskStatus"`
	Comment        string    `json:"comment" bson:"comment"`
}

func CreateDlTask(id, dataset, loss string, epochs, batchSize int, nnArchitecture []NnLayer) (*DlTask, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	task := DlTask{
		TaskId:         id,
		Dataset:        dataset,
		Loss:           loss,
		Epochs:         epochs,
		BatchSize:      batchSize,
		NnArchitecture: nnArchitecture,
		CreateTime:     jtime.NowStr(),
		TaskStatus:     "执行中",
	}
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteDlTask(id string) (int64, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	ret, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return ret.DeletedCount, nil
}

func ListDlTask() ([]DlTask, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
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
	records := make([]DlTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), mongoCtxTimeout)
	for cur.Next(ctx) {
		result := DlTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetDlTaskData(id string) (bson.A, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "data": true} // 注意 _id 默认会返回，需要手动过滤
	var result bson.M
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["data"].(bson.A), nil
}
