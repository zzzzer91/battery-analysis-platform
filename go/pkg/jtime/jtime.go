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
	t.Time = tTmp
	return nil
}

// UnmarshalBSON 实现 Unmarshaler 接口，支持自定义类型的 mongo 反序列化
func (t *Time) UnmarshalBSON(b []byte) error {
	// b 字节流有个 4 字节头部，值为 20，可能代表了类型
	// 还有个 1 字节的尾部，0
	tTmp, err := time.ParseInLocation(FormatLayout, string(b[4:len(b)-1]), time.Local)
	if err != nil {
		return err
	}
	t.Time = tTmp
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
