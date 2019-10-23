package model

import (
	"battery-analysis-platform/app/main/db"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	mongoCtxTimeout = time.Second * 5
)

const (
	mongoCollectionMiningTasks = "mining_tasks"
	mongoCollectionMlTasks     = "ml_tasks"
)

// 确保创建 mongo 索引
func createMongoCollectionIdx(name string, model mongo.IndexModel) error {
	collection := db.Mongo.Collection(name)
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.Indexes().CreateOne(
		ctx,
		model,
	)
	return err
}

func init() {
	model := mongo.IndexModel{
		Keys: bson.M{
			"taskId": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionMiningTasks, model); err != nil {
		panic(err)
	}
	if err := createMongoCollectionIdx(mongoCollectionMlTasks, model); err != nil {
		panic(err)
	}
}
