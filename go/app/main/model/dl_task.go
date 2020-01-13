package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/conv"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NnLayer struct {
	Neurons    int    `json:"neurons" bson:"neurons"`
	Activation string `json:"activation" bson:"activation"`
}

type NnHyperParameter struct {
	HiddenLayerStructure  []NnLayer `json:"hiddenLayerStructure" bson:"hiddenLayerStructure"`
	OutputLayerActivation string    `json:"outputLayerActivation" bson:"outputLayerActivation"`
	Loss                  string    `json:"loss" bson:"loss"`
	Seed                  int       `json:"seed" bson:"seed"`
	BatchSize             int       `json:"batchSize" bson:"batchSize"`
	Epochs                int       `json:"epochs" bson:"epochs"`
	LearningRate          float64   `json:"learningRate" bson:"learningRate"`
}

type NnTrainingHistory struct {
	Loss     []float64 `json:"loss" bson:"loss"`
	Accuracy []float64 `json:"accuracy" bson:"accuracy"`
}

type NnEvalResult struct {
	A1Count     int `json:"a1Count" bson:"a1Count"`
	A2Count     int `json:"a2Count" bson:"a2Count"`
	A3Count     int `json:"a3Count" bson:"a3Count"`
	A4Count     int `json:"a4Count" bson:"a4Count"`
	AOtherCount int `json:"aOtherCount" bson:"aOtherCount"`
}

type DlTask struct {
	BaseTask        `bson:",inline"`
	Dataset         string             `json:"dataset" bson:"dataset"`
	HyperParameter  *NnHyperParameter  `json:"hyperParameter" bson:"hyperParameter"`
	TrainingHistory *NnTrainingHistory `json:"-" bson:"trainingHistory"`
	EvalResult      *NnEvalResult      `json:"-" bson:"evalResult"`
}

func CreateDlTask(id, dataset string, hyperParameter *NnHyperParameter) (*DlTask, error) {
	task := DlTask{
		BaseTask:       newBaseTask(id),
		Dataset:        dataset,
		HyperParameter: hyperParameter,
	}
	err := creatTask(mongoCollectionDlTask, task)

	return &task, err
}

func DeleteDlTask(id string) error {
	return deleteTask(mongoCollectionDlTask, id)
}

func ListDlTask() ([]DlTask, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{}                                                  // 过滤记录
	projection := bson.M{"trainingHistory": false, "evalResult": false} // 过滤字段
	sort := bson.M{"createTime": -1}                                    // 结果排序
	// 注意 ctx 不能几个连接复用
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]DlTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), mongoCtxTimeout)
	for cur.Next(ctx) {
		result := DlTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetDlTaskTrainingHistory(id string, readFromRedis bool) (*NnTrainingHistory, error) {
	if readFromRedis {
		lossStrList, err := db.Redis.LRange(
			"deeplearningTask:trainingHistory:"+id+":loss",
			0, -1).Result()
		if err != nil {
			return nil, err
		}
		accuracyStrList, err := db.Redis.LRange(
			"deeplearningTask:trainingHistory:"+id+":accuracy",
			0, -1).Result()
		if err != nil {
			return nil, err
		}

		// 转换为 float
		lossList, err := conv.StringSlice2FloatSlice(lossStrList)
		if err != nil {
			return nil, err
		}
		accuracyList, err := conv.StringSlice2FloatSlice(accuracyStrList)
		if err != nil {
			return nil, err
		}

		return &NnTrainingHistory{
			Loss:     lossList,
			Accuracy: accuracyList,
		}, nil
	} else {
		collection := db.Mongo.Collection(mongoCollectionDlTask)
		filter := bson.M{"taskId": id}
		projection := bson.M{"_id": false, "trainingHistory": true}
		var result DlTask
		ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
		err := collection.FindOne(ctx, filter, options.FindOne().
			SetProjection(projection)).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result.TrainingHistory, nil
	}
}

func GetDlTaskEvalResult(id string) (*NnEvalResult, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "evalResult": true}
	var result DlTask
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.EvalResult, nil
}
