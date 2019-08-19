// http://axiaoxin.com/article/241/
// 使 json序列化时间时，输出指定格式
// 我们需定义一个内嵌time.Time的结构体，并重写MarshalJSON方法，
// 然后在定义model的时候把time.Time类型替换为我们自己的类型即可。
// 但是在gorm中只重写MarshalJSON是不够的，
// 只写这个方法会在写数据库的时候会提示delete_at字段不存在，
// 需要加上database/sql的Value和Scan方法

package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		s := fmt.Sprintf("can not convert %v to timestamp", v)
		panic(s)
	}
	*t = JSONTime{Time: value}
	return nil
}

func Now() *JSONTime {
	t := time.Now()
	return &JSONTime{Time: t}
}
