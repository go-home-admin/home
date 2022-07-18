package database

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func Now() Time {
	return Time{
		time.Now(),
	}
}

func NowPointer() *Time {
	return &Time{
		time.Now(),
	}
}

// Value 写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		t.Time = value
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	*t = StrToTime(strings.Replace(string(data), `"`, "", -1))
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	data := `"` + t.YmdHis() + `"`
	return []byte(data), nil
}

func (t Time) YmdHis() string {
	return t.Format("2006-01-02 15:04:05")
}

func (t Time) Ymd() string {
	return t.Format("2006-01-02")
}

func (t Time) DayEnd() string {
	return t.Format("2006-01-02") + " 23:59:59"
}

// StrToTime 字符串转时间类型，仅支持两种格式
func StrToTime(str string) Time {
	tm, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		tm, err = time.ParseInLocation("2006-01-02", str, time.Local)
	}
	return Time{tm}
}

// UnixToTime 时间戳转时间类型
func UnixToTime(unix int64) Time {
	return Time{time.Unix(unix, 0)}
}

// TimeToIntDate 时间转为INT的日期格式（20060102）
func (t Time) TimeToIntDate() int32 {
	v, _ := strconv.ParseInt(t.Format("20060102"), 10, 32)
	return int32(v)
}

// IntDateToTime INT的日期格式转为时间
func IntDateToTime(date int32) Time {
	tm, _ := time.ParseInLocation("20060102", fmt.Sprint(date), time.Local)
	return Time{tm}
}
