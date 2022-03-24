package database

import "time"

type Time time.Time

func (t Time) YmdHis() string {
	return ""
}

type Json string
