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
	mongoCollectionYuTongVehicle = "yutong_vehicle"
	mongoCollectionBeiQiVehicle  = "beiqi_vehicle"
	mongoCollectionMiningTask    = "mining_task"
	mongoCollectionDlTask        = "deeplearning_task"
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
	// yutong_vehicle
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"状态号": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}

	// beiqi_vehicle
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"动力电池充放电状态": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}

	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"taskId": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := createMongoCollectionIdx(mongoCollectionMiningTask, indexModel); err != nil {
		panic(err)
	}
	if err := createMongoCollectionIdx(mongoCollectionDlTask, indexModel); err != nil {
		panic(err)
	}
}
