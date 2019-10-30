package model

import (
	"battery-analysis-platform/app/main/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

const (
	TaskStatusPreparing  = 0
	TaskStatusProcessing = 1
	TaskStatusSuccess    = 6
	TaskStatusFailure    = 7
)

const (
	TaskWaitSigTimeout = time.Second * 30
)

// 注意：
// 结构体嵌入其他结构体时，要指定 `bson:",inline"`（注意有个逗号）标签才能正确读入 bson，
// 同时该结构体要public；
// 而用 json 序列化时，直接嵌入结构体即可，不需要做其他事。
// 嵌入时，不能用指针
type BaseTask struct {
	TaskId     string `json:"taskId" bson:"taskId"`
	CreateTime string `json:"createTime" bson:"createTime"`
	TaskStatus int    `json:"taskStatus" bson:"taskStatus"`
	Comment    string `json:"comment" bson:"comment"`
}

func creatTask(collectionName string, task interface{}) error {
	collection := db.Mongo.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.InsertOne(ctx, task)
	return err
}

func deleteTask(collectionName string, id string) error {
	collection := db.Mongo.Collection(collectionName)
	filter := bson.M{"taskId": id}
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
