package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
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

type DlTask struct {
	BaseTask       `bson:",inline"`
	Dataset        string            `json:"dataset" bson:"dataset"`
	HyperParameter *NnHyperParameter `json:"hyperParameter" bson:"hyperParameter"`
}

func CreateDlTask(id, dataset string, hyperParameter *NnHyperParameter) (*DlTask, error) {
	task := DlTask{
		BaseTask: BaseTask{
			TaskId:     id,
			CreateTime: jtime.NowStr(),
			TaskStatus: "执行中",
		},
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
	filter := bson.M{}                             // 过滤记录
	projection := bson.M{"trainingHistory": false} // 过滤字段
	sort := bson.M{"createTime": -1}               // 结果排序
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

func GetDlTaskTrainingHistory(id string) (bson.M, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "trainingHistory": true}
	var result bson.M
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["trainingHistory"].(bson.M), nil
}

func GetDlTaskEvalResult(id string) (bson.M, error) {
	collection := db.Mongo.Collection(mongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "evalResult": true}
	var result bson.M
	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["evalResult"].(bson.M), nil
}
