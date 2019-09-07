package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

var MongoDB *mongo.Database

func InitMongo(uri string) {
	index := strings.LastIndex(uri, "/")
	uriSplit := uri[:index]
	database := uri[index+1:]

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uriSplit))
	if err != nil {
		panic(err)
	}

	MongoDB = client.Database(database)
}
