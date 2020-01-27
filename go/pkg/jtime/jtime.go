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
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"time"
)

const (
	FormatLayout = "2006-01-02 15:04:05"
)

type Time struct {
	time.Time
}

// MarshalJSON on Time format Time field with %Y-%m-%d %H:%M:%S
func (t Time) MarshalJSON() ([]byte, error) {
	formatted := `"` + t.Format(FormatLayout) + `"`
	return []byte(formatted), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	s := string(data)
	// Ignore null, like in the main JSON package.
	if s == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	t.Time, err = time.Parse(`"`+FormatLayout+`"`, s)
	return err
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	tTmp, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	t.Time = tTmp.Local()
	return nil
}

// MarshalBSON 有问题，bson.Marshal 不能对结构体成员使用
// 无法获取成员对应的 bson 标签
//func (t Time) MarshalBSON() ([]byte, error) {
//	_, a, b := bson.MarshalValue(t.Time)
//	return a, b
//}

func (t Time) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(t.Time)
}

// UnmarshalBSON 实现 Unmarshaler 接口，支持自定义类型的 mongo 反序列化
func (t *Time) UnmarshalBSON(b []byte) error {
	tTmep, _, ok := bsoncore.ReadTime(b)
	if !ok {
		return errors.New("jtime.Time UnmarshalBSON error")
	}
	t.Time = tTmep
	return nil
}

func Wrap(t time.Time) Time {
	return Time{
		Time: t,
	}
}

func Now() Time {
	return Wrap(time.Now())
}

func NowStr() string {
	return time.Now().Format(FormatLayout)
}
