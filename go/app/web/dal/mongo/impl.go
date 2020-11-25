package mongo

import (
	"battery-analysis-platform/app/web/constant"
	"battery-analysis-platform/app/web/model"
	"battery-analysis-platform/pkg/jtime"
	"battery-analysis-platform/pkg/security"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type serviceImpl struct {
	cli *mongo.Database
}

func (s *serviceImpl) creatTask(collectionName string, task interface{}) error {
	return s.insertMongoCollection(collectionName, task)
}

func (s *serviceImpl) deleteTask(collectionName string, id string) error {
	collection := s.cli.Collection(collectionName)
	filter := bson.D{{"taskId", id}}
	ctx := newTimeoutCtx()
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

// 确保创建 mongo 索引
func (s *serviceImpl) createMongoCollectionIdx(name string, model mongo.IndexModel) error {
	collection := s.cli.Collection(name)
	ctx := newTimeoutCtx()
	_, err := collection.Indexes().CreateOne(ctx, model)
	return err
}

// 在 collection 中插入一条记录
func (s *serviceImpl) insertMongoCollection(collectionName string, item interface{}) error {
	collection := s.cli.Collection(collectionName)
	ctx := newTimeoutCtx()
	_, err := collection.InsertOne(ctx, item)
	return err
}

func (s *serviceImpl) init() {
	// user
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionUser, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"type": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionUser, indexModel); err != nil {
		panic(err)
	}

	// yutong_vehicle
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"状态号": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionYuTongVehicle, indexModel); err != nil {
		panic(err)
	}

	// beiqi_vehicle
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"时间": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"状态号": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionBeiQiVehicle, indexModel); err != nil {
		panic(err)
	}

	// task
	indexModel = mongo.IndexModel{
		Keys: bson.M{
			"taskId": 1,
		},
		Options: options.Index().SetUnique(false),
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionMiningTask, indexModel); err != nil {
		panic(err)
	}
	if err := s.createMongoCollectionIdx(constant.MongoCollectionDlTask, indexModel); err != nil {
		panic(err)
	}
}

func (s *serviceImpl) GetBatteryList(tableName string, startDate time.Time, limit int, fields []string) ([]bson.M, error) {
	collection := s.cli.Collection(tableName)
	filter := bson.D{{"时间", bson.D{{"$gte", startDate}}}}
	// 注意 _id 默认会返回，需要手动过滤
	projection := bson.D{
		{"_id", false},
		{"时间", true},
		{"状态号", true},
	}
	for _, field := range fields {
		projection = append(projection, bson.E{Key: field, Value: true})
	}

	ctx := newTimeoutCtx()

	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]bson.M, 0)
	for cur.Next(ctx) {
		result := make(bson.M)
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		// 为了使 json 序列化时得到想要格式
		temp := result["时间"].(primitive.DateTime)
		result["时间"] = jtime.Wrap(temp.Time())
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func (s *serviceImpl) CreateDlTask(id, dataset string, hyperParameter *model.NnHyperParameter) (*model.DlTask, error) {
	task := model.NewDlTask(id, dataset, hyperParameter)
	err := s.creatTask(constant.MongoCollectionDlTask, task)

	return task, err
}

func (s *serviceImpl) GetDlTaskList() ([]model.DlTask, error) {
	collection := s.cli.Collection(constant.MongoCollectionDlTask)
	filter := bson.D{}
	projection := bson.D{
		{"_id", false},
		{"trainingHistory", false},
		{"evalResult", false},
	}
	sort := bson.D{{"createTime", -1}}
	// 注意 ctx 不能几个连接复用
	ctx, _ := context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.DlTask, 0)
	ctx, _ = context.WithTimeout(context.Background(), constant.MongoCtxTimeout)
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

func (s *serviceImpl) GetDlTaskTrainingHistory(id string) (*model.NnTrainingHistory, error) {
	collection := s.cli.Collection(constant.MongoCollectionDlTask)
	filter := bson.D{{"taskId", id}}
	projection := bson.D{{"_id", false}, {"trainingHistory", true}}
	var result model.DlTask
	ctx := newTimeoutCtx()
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.TrainingHistory, nil
}

func (s *serviceImpl) GetDlTaskEvalResult(id string) (*model.NnEvalResult, error) {
	collection := s.cli.Collection(constant.MongoCollectionDlTask)
	filter := bson.D{{"taskId", id}}
	projection := bson.D{{"_id", false}, {"evalResult", true}}
	var result model.DlTask
	ctx := newTimeoutCtx()
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.EvalResult, nil
}

func (s *serviceImpl) DeleteDlTask(id string) error {
	return s.deleteTask(constant.MongoCollectionDlTask, id)
}

func (s *serviceImpl) CreateUser(name, password, comment string) (*model.User, error) {
	user := model.NewUser(name, password, comment)
	err := s.insertMongoCollection(constant.MongoCollectionUser, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *serviceImpl) GetCommonUserList() ([]model.User, error) {
	collection := s.cli.Collection(constant.MongoCollectionUser)
	filter := bson.D{{"type", bson.D{{"$ne", constant.UserTypeSuperUser}}}} // 过滤记录
	projection := bson.D{{"_id", false}}
	ctx := newTimeoutCtx()
	cur, err := collection.Find(ctx, filter, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	users := make([]model.User, 0)
	for cur.Next(ctx) {
		result := model.User{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, result)
	}
	_ = cur.Close(ctx)
	return users, nil
}

func (s *serviceImpl) GetUser(name string) (*model.User, error) {
	var user model.User
	collection := s.cli.Collection(constant.MongoCollectionUser)
	filter := bson.D{{"name", name}}
	projection := bson.D{{"_id", false}} // 注意 _id 默认会返回，需要手动过滤
	ctx := newTimeoutCtx()
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *serviceImpl) UpdateUserInfo(user *model.User) error {
	collection := s.cli.Collection(constant.MongoCollectionUser)
	filter := bson.D{{"name", user.Name}} // 过滤记录
	update := bson.D{{"$set", bson.D{
		{"comment", user.Comment},
		{"status", user.Status},
	}}}
	ctx := newTimeoutCtx()
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *serviceImpl) UpdateUserLoginTimeAndCount(user *model.User) error {
	collection := s.cli.Collection(constant.MongoCollectionUser)
	filter := bson.D{{"name", user.Name}} // 过滤记录
	update := bson.D{{"$set", bson.D{
		{"lastLoginTime", user.LastLoginTime},
		{"loginCount", user.LoginCount},
	}}}
	ctx := newTimeoutCtx()
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *serviceImpl) UpdateUserPassword(userName, password string) error {
	collection := s.cli.Collection(constant.MongoCollectionUser)
	filter := bson.D{{"name", userName}} // 过滤记录
	ph, err := security.GeneratePasswordHash(password)
	if err != nil {
		return err
	}
	update := bson.D{{"$set", bson.D{
		{"password", ph},
	}}}
	ctx := newTimeoutCtx()
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}

func (s *serviceImpl) CreateMiningTask(id, name, dataComeFrom, dateRange string) (*model.MiningTask, error) {
	task := model.NewMiningTask(id, name, dataComeFrom, dateRange)
	err := s.creatTask(constant.MongoCollectionMiningTask, task)
	return task, err
}

func (s *serviceImpl) GetMiningTaskList() ([]model.MiningTask, error) {
	collection := s.cli.Collection(constant.MongoCollectionMiningTask)
	filter := bson.D{}
	projection := bson.D{{"_id", false}, {"data", false}}
	sort := bson.D{{"createTime", -1}}
	// 注意 ctx 不能几个连接复用，原因见 `context.WithTimeout` 源码
	ctx := newTimeoutCtx()
	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		return nil, err
	}
	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]model.MiningTask, 0)
	for cur.Next(ctx) {
		result := model.MiningTask{}
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}

func (s *serviceImpl) GetMiningTaskData(id string) (bson.A, error) {
	collection := s.cli.Collection(constant.MongoCollectionMiningTask)
	filter := bson.D{{"taskId", id}}
	projection := bson.D{{"_id", false}, {"data", true}}
	ctx := newTimeoutCtx()
	var result bson.M
	err := collection.FindOne(ctx, filter,
		options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result["data"].(bson.A), nil
}

func (s *serviceImpl) DeleteMiningTask(id string) error {
	return s.deleteTask(constant.MongoCollectionMiningTask, id)
}
