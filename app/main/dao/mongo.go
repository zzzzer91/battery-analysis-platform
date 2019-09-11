package dao

import (
	"battery-anlysis-platform/pkg/conf"
	"battery-anlysis-platform/pkg/db"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoDB *mongo.Database

func init() {
	d, err := db.InitMongo(&conf.App.Main.Mongo)
	if err != nil {
		panic(err)
	}
	MongoDB = d
}
