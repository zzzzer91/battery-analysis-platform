package model

import (
	"battery-analysis-platform/app/main/db"
	"battery-analysis-platform/pkg/jtime"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type dbTable struct {
	Name     string
	FieldSet map[string]struct{}
}

var BatteryNameToTable map[string]dbTable

func init() {
	yutongFieldSet := map[string]struct{}{
		"时间":      {},
		"总电压":     {},
		"总电流":     {},
		"车速":      {},
		"正向累计电量":  {},
		"反向累计电量":  {},
		"总里程":     {},
		"SOC":     {},
		"单体最高温度":  {},
		"单体最低温度":  {},
		"单体最高电压":  {},
		"单体最低电压":  {},
		"最高温度电池号": {},
		"最低温度电池号": {},
		"最高电压电池号": {},
		"最低电压电池号": {},
	}

	beiqiFieldSet := map[string]struct{}{
		"时间":          {},
		"动力电池内部总电压V1": {},
		"动力电池充/放电电流":  {},
		"动力电池可用能量":    {},
		"动力电池可用容量":    {},
		"动力电池剩余电量SOC": {},
		"MSODO总里程":    {},
	}

	BatteryNameToTable = map[string]dbTable{
		"宇通_4F37195C1A908CFBE0532932A8C0EECB": {
			Name: mongoCollectionYuTongVehicle, FieldSet: yutongFieldSet,
		},
		"北汽_LNBSCU3HXJR884327": {
			Name: mongoCollectionBeiQiVehicle, FieldSet: beiqiFieldSet,
		},
	}
}

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

	ctx, _ := context.WithTimeout(context.Background(), mongoCtxTimeout)
	cur, err := collection.Find(ctx, filter,
		options.Find().SetProjection(projection).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	// 修复时区，go 默认 UTC，用时间区间查询会出问题
	cstZone := time.FixedZone("CST", 8*3600)

	// 为了使其找不到时返回空列表，而不是 nil
	records := make([]bson.M, 0)
	ctx, _ = context.WithTimeout(context.Background(), mongoCtxTimeout)
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
