package database

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type PgTimeTz struct {
	time.Time
}

// Value Implements Valuer interface for writing to DB
func (pt PgTimeTz) Value() (driver.Value, error) {
	// 格式化为 HH:MM:SS.ssssss
	return pt.String(), nil
}

// Scan Implements Scanner interface for reading from DB
func (pt *PgTimeTz) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		pt.Time = v
	case []byte:
		// Parse byte array in HH:MM:SS[.ssssss]+TZ format
		t, err := time.Parse("15:04:05.999999-07", string(v))
		if err != nil {
			return fmt.Errorf("timetz parsing failed: %v", err)
		}
		pt.Time = t
	case string:
		t, err := time.Parse("15:04:05.999999-07", v)
		if err != nil {
			return fmt.Errorf("timetz parsing failed: %v", err)
		}
		pt.Time = t
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

func (pt PgTimeTz) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, pt.Format("15:04:05.999999-07"))), nil
}

func (pt *PgTimeTz) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"15:04:05.999999-07"`, string(data))
	if err != nil {
		return err
	}
	pt.Time = t
	return nil
}

func (pt PgTimeTz) String() string {
	return pt.Format("15:04:05.999999-07")
}

func (pt PgTimeTz) His(loc *time.Location) string {
	return pt.Time.In(loc).Format("15:04:05")
}

func (pt PgTimeTz) Add(d time.Duration) PgTime {
	return PgTime{
		Time: pt.Time.Add(d),
	}
}

func StrToPgTimeTz(str string) PgTimeTz {
	tm, _ := time.Parse("15:04:05.999999-07", str)
	return PgTimeTz{tm}
}
