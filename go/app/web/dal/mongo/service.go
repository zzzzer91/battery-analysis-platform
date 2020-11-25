package mongo

import (
	"battery-analysis-platform/app/web/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Service interface {
	GetBatteryList(tableName string, startDate time.Time, limit int, fields []string) ([]bson.M, error)
	CreateDlTask(id, dataset string, hyperParameter *model.NnHyperParameter) (*model.DlTask, error)
	GetDlTaskList() ([]model.DlTask, error)
	GetDlTaskTrainingHistory(id string) (*model.NnTrainingHistory, error)
	GetDlTaskEvalResult(id string) (*model.NnEvalResult, error)
	DeleteDlTask(id string) error
	CreateUser(name, password, comment string) (*model.User, error)
	GetCommonUserList() ([]model.User, error)
	GetUser(name string) (*model.User, error)
	UpdateUserInfo(user *model.User) error
	UpdateUserLoginTimeAndCount(user *model.User) error
	UpdateUserPassword(userName, password string) error
	CreateMiningTask(id, name, dataComeFrom, dateRange string) (*model.MiningTask, error)
	GetMiningTaskList() ([]model.MiningTask, error)
	GetMiningTaskData(id string) (bson.A, error)
	DeleteMiningTask(id string) error
}

func NewService(cli *mongo.Database) Service {
	s := &serviceImpl{cli}
	s.init()
	return s
}
