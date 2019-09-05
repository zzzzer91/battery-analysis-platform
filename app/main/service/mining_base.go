package service

import (
	"battery-anlysis-platform/app/main/dao"
	"battery-anlysis-platform/pkg/mysqlx"
	"errors"
	"strings"
)

type MiningBaseService struct {
	TableName string
	StartDate string
	Fields    []string
	DataLimit int
}

func (s *MiningBaseService) Query() ([]map[string]interface{}, error) {
	rows, err := dao.MysqlDB.Table(s.TableName).
		Where("timestamp >= ?", s.StartDate).
		Select("timestamp," + strings.Join(s.Fields, ",")).
		Limit(s.DataLimit).
		Rows()
	if err != nil {
		return nil, errors.New("查询失败")
	}

	records, err := mysqlx.GetRecords(rows)
	if err != nil {
		panic(err)
	}
	if len(records) == 0 {
		return nil, errors.New("未查询到相关数据")
	}
	return records, nil
}
