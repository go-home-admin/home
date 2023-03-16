package database

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
)

type JSON []byte

func NewJSON(v interface{}) JSON {
	str, err := json.Marshal(v)
	if err != nil {
		log.Errorf("NewJSON err %v, arg %v", err, v)
		return []byte("")
	}
	return str
}

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		st, ok := value.(string)
		if !ok {
			log.Warnf("json scan value error, %v; 只支持[]byte|string", value)
		} else {
			s = []byte(st)
		}
	}
	*j = append((*j)[0:0], s...)
	return nil
}

func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

func (j JSON) Trans(v interface{}) error {
	b, _ := j.Value()
	err := json.Unmarshal([]byte(b.(string)), &v)
	return err
}
