// 这个包封装了使用原生 sql 对 mysql 不定字段查询的一些函数

package mysqlx

import (
	"battery-analysis-platform/pkg/jtime"
	"database/sql"
	"reflect"
)

// 这些类型接收 mysql 返回的数据，因为可能为 NULL，所以要用指针
var (
	typeFloat32 = reflect.TypeOf((*float32)(nil))
	typeFloat64 = reflect.TypeOf((*float64)(nil))
	typeInt     = reflect.TypeOf((*int)(nil))
	typeString  = reflect.TypeOf((*string)(nil))
	// 自定义时间类型，用于 json 序列化时格式化时间
	typeJSONTime = reflect.TypeOf((*jtime.JSONTime)(nil))
)

func getGoType(databaseTypeName string) (tp reflect.Type) {
	switch databaseTypeName {
	case "VARCHAR", "TEXT", "DATE", "TIME":
		tp = typeString
	case "INT", "TINYINT":
		tp = typeInt
	case "DOUBLE":
		tp = typeFloat64
	case "FLOAT":
		tp = typeFloat32
	case "DECIMAL":
		tp = typeFloat64
	case "TIMESTAMP", "DATETIME":
		tp = typeJSONTime
	}
	return
}

// 将 mysql-driver 返回的行转换成 []map[string]interface{}，用于 json 序列化
func GetRecords(rows *sql.Rows) ([]map[string]interface{}, error) {
	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	// 获取字段名和字段类型，字段类型通过反射
	columnNames := make([]string, 0, len(columns))
	columnTypes := make([]reflect.Type, 0, len(columns))
	for _, column := range columns {
		columnNames = append(columnNames, column.Name())
		columnTypes = append(columnTypes, getGoType(column.DatabaseTypeName()))
	}

	var records []map[string]interface{}
	scanArgs := make([]interface{}, len(columns))
	for rows.Next() {
		record := make(map[string]interface{}, len(columnNames))
		for i := range columnNames {
			// 运用反射，在运行时创建变量
			tmp := reflect.New(columnTypes[i]).Interface()
			record[columnNames[i]] = tmp
			scanArgs[i] = tmp
		}

		err := rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}
