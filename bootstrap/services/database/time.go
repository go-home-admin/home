package database

import (
	"database/sql/driver"
	"fmt"
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

func (t *Time) YmdHis() string {
	return t.Format("2006-01-02 15:04:05")
}
