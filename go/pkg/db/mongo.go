package db

import (
	"battery-analysis-platform/pkg/conf"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func InitMongo(mongoConf *conf.MongoConf) (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConf.Uri))
	if err != nil {
		return nil, err
	}

	return client.Database(mongoConf.Database), nil
}
