package dao

import (
	"battery-analysis-platform/app/main/conf"
	"battery-analysis-platform/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDB *mongo.Database

func init() {
	d, err := db.InitMongo(&conf.App.Mongo)
	if err != nil {
		panic(err)
	}
	MongoDB = d
}
