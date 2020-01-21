package dao

import (
	"battery-analysis-platform/app/main/consts"
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/app/main/model"
	"battery-analysis-platform/pkg/conv"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDlTask(id, dataset string, hyperParameter *model.NnHyperParameter) (*model.DlTask, error) {
	task := model.NewDlTask(id, dataset, hyperParameter)
	err := creatTask(consts.MongoCollectionDlTask, task)

	return task, err
}

func DeleteDlTask(id string) error {
	return deleteTask(consts.MongoCollectionDlTask, id)
}

func ListDlTask() ([]model.DlTask, error) {
	collection := db.Mongo.Collection(consts.MongoCollectionDlTask)
	filter := bson.M{}                                                  // 过滤记录
	projection := bson.M{"trainingHistory": false, "evalResult": false} // 过滤字段
	sort := bson.M{"createTime": -1}                                    // 结果排序
	// 注意 ctx 不能几个连接复用
	ctx, _ := context.WithTimeout(context.Background(), consts.MongoCtxTimeout)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.DlTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), consts.MongoCtxTimeout)
	for cur.Next(ctx) {
		result := model.DlTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func GetDlTaskTrainingHistory(id string, readFromRedis bool) (*model.NnTrainingHistory, error) {
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

		return &model.NnTrainingHistory{
			Loss:     lossList,
			Accuracy: accuracyList,
		}, nil
	} else {
		collection := db.Mongo.Collection(consts.MongoCollectionDlTask)
		filter := bson.M{"taskId": id}
		projection := bson.M{"_id": false, "trainingHistory": true}
		var result model.DlTask
		ctx, _ := context.WithTimeout(context.Background(), consts.MongoCtxTimeout)
		err := collection.FindOne(ctx, filter, options.FindOne().
			SetProjection(projection)).Decode(&result)
		if err != nil {
			return nil, err
		}
		return result.TrainingHistory, nil
	}
}

func GetDlTaskEvalResult(id string) (*model.NnEvalResult, error) {
	collection := db.Mongo.Collection(consts.MongoCollectionDlTask)
	filter := bson.M{"taskId": id}
	projection := bson.M{"_id": false, "evalResult": true}
	var result model.DlTask
	ctx, _ := context.WithTimeout(context.Background(), consts.MongoCtxTimeout)
	err := collection.FindOne(ctx, filter, options.FindOne().
		SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.EvalResult, nil
}
