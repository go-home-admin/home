package database

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type PgTime struct {
	time.Time
}

// Value Implements Valuer interface for writing to DB
func (pt PgTime) Value() (driver.Value, error) {
	// 格式化为 HH:MM:SS.ssssss
	return pt.String(), nil
}

// Scan Implements Scanner interface for reading from DB
func (pt *PgTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		// Extract time component (date set to 0000-01-01)
		pt.Time = time.Date(0, 1, 1, v.Hour(), v.Minute(), v.Second(), v.Nanosecond(), time.UTC)
	case []byte:
		// Parse byte array in HH:MM:SS[.ssssss] format
		t, err := time.Parse("15:04:05.999999", string(v))
		if err != nil {
			return fmt.Errorf("time parsing failed: %v", err)
		}
		pt.Time = t
	case string:
		t, err := time.Parse("15:04:05.999999", v)
		if err != nil {
			return fmt.Errorf("time parsing failed: %v", err)
		}
		pt.Time = t
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

func (pt PgTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, pt.String())), nil
}

func (pt *PgTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"15:04:05.999999"`, string(data))
	if err != nil {
		return err
	}
	pt.Time = t
	return nil
}

func (pt PgTime) String() string {
	return pt.Format("15:04:05.999999")
}

func (pt PgTime) His() string {
	return pt.Format("15:04:05")
}

func (pt PgTime) Add(d time.Duration) PgTime {
	return PgTime{
		Time: pt.Time.Add(d),
	}
}

func StrToPgTime(str string) PgTime {
	tm, _ := time.Parse("15:04:05.999999", str)
	return PgTime{tm}
}
