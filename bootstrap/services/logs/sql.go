package logs

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strings"
	"time"
)

type SqlLog struct {
	L *log.Logger
	logger.Config

	infoStr, warnStr, errStr            string
	traceStr, traceErrStr, traceWarnStr string
	// 只要存在这个字符串就跳过
	skips []string
}

// NewSqlLog skip 需要屏蔽的sql
func NewSqlLog(l *log.Logger, config logger.Config, skips ...string) logger.Interface {
	var (
		infoStr      = "\n[info] "
		warnStr      = "\n[warn] "
		errStr       = "\n[error] "
		traceStr     = "\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
	)

	got := &SqlLog{
		L:      l,
		Config: config,

		infoStr:      infoStr,
		warnStr:      warnStr,
		errStr:       errStr,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,

		skips: skips,
	}

	return got
}

// LogMode log mode
func (l *SqlLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}
func (l *SqlLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.L.Info(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}
}
func (l *SqlLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.L.Warn(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}
}
func (l *SqlLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.L.Error(msg, append([]interface{}{utils.FileWithLineNum()}, data...))
	}
}

func (l *SqlLog) Printf(format string, args ...interface{}) {
	l.L.Printf(format, args...)
}

func (l *SqlLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}
	switch {
	case err != nil && !errors.Is(err, logger.ErrRecordNotFound):
		sql, _ := fc()
		l.L.WithFields(log.Fields{"sql": sql, "file": utils.FileWithLineNum()}).Error(err.Error())
	case l.LogLevel >= logger.Info:
		elapsed := time.Since(begin)
		sql, rows := fc()

		for _, skip := range l.skips {
			if strings.Index(sql, skip) != -1 {
				return
			}
		}
		l.L.WithFields(log.Fields{"type": "query", "begin": begin, "row": rows, "t": float64(elapsed.Nanoseconds()) / 1e6}).Debug(sql)
	}
}
