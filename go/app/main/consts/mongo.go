package consts

import "time"

const (
	MongoCtxTimeout = time.Second * 5
)

const (
	MongoCollectionUser          = "user"
	MongoCollectionYuTongVehicle = "yutong_vehicle"
	MongoCollectionBeiQiVehicle  = "beiqi_vehicle"
	MongoCollectionMiningTask    = "mining_task"
	MongoCollectionDlTask        = "deeplearning_task"
)
