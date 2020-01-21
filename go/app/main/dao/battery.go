package dao

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func GetBatteryData(tableName, startDate string, limit int, fields []string) ([]bson.M, error) {
	collection := db.Mongo.Collection(tableName)

	// 查询指定范围时间的数据
	sDate, err := time.ParseInLocation(jtime.FormatLayout, startDate, time.Local)
	if err != nil {
		return nil, err
	}
	// filter := bson.M{"时间": bson.M{"$gte": sDate, "$lt": eDate}}
	filter := bson.M{"时间": bson.M{"$gte": sDate}}

	projection := bson.M{"_id": false, "时间": true, "状态号": true}
	for _, field := range fields {
		projection[field] = true
	}

	ctx := NewTimeoutCtx()

	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	// 修复时区，go 默认 UTC，用时间区间查询会出问题
	cstZone := time.FixedZone("CST", 8*3600)

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
		result["时间"] = jtime.Wrap(temp.Time().In(cstZone))
		records = append(records, result)
	}
	_ = cur.Close(ctx)
	return records, nil
}
