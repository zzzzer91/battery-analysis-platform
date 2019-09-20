package model

import (
	"battery-analysis-platform/app/main/dao"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	dao.MysqlDB.AutoMigrate(&User{})

	// 确保 mongo 索引
	collection := dao.MongoDB.Collection(mongoCollectionMiningTasks)
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	_, err := collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.M{
				"taskId": 1,
			},
			Options: options.Index().SetUnique(false),
		},
	)
	if err != nil {
		panic(err)
	}
}
