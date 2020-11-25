package dal

import (
	"battery-analysis-platform/app/web/conf"
	"battery-analysis-platform/app/web/dal/mongo"
	"battery-analysis-platform/pkg/db"
)

var mongoService mongo.Service

func GetMongoService() mongo.Service {
	return mongoService
}

func Init() {
	cli, err := db.InitMongo(&conf.App.Mongo)
	if err != nil {
		panic(err)
	}
	mongoService = mongo.NewService(cli)
}
