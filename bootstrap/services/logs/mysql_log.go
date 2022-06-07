package logs

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

// MysqlLog 调试的过程
type MysqlLog struct{}

func (d *MysqlLog) LogMode(level logger.LogLevel) logger.Interface {
	return newDebugLog()
}

func (d *MysqlLog) Info(ctx context.Context, s string, i ...interface{}) {
	log.Info(s, i)
}

func (d *MysqlLog) Warn(ctx context.Context, s string, i ...interface{}) {
	log.Warn(s, i)
}

func (d *MysqlLog) Error(ctx context.Context, s string, i ...interface{}) {
	log.Error(s, i)
}

func (d *MysqlLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, rows := fc()
	log.WithFields(log.Fields{"path": "sql", "begin": begin, "row": rows}).Debug(sql)
}

func newDebugLog() logger.Interface {
	return &MysqlLog{}
}
