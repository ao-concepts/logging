package logging

import (
	"context"
	"time"

	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	log Logger
}

func (l *gormLogger) LogMode(glogger.LogLevel) glogger.Interface {
	return l // do nothing here
}

func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.log.GetLevel() <= Info {
		l.log.Info(msg, data...)
	}
}

func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.log.GetLevel() <= Warn {
		l.log.Warn(msg, data...)
	}
}

func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.log.GetLevel() <= Error {
		l.log.Error(msg, data...)
	}
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if err != nil {
		l.log.ErrError(err)
		return
	}

	level := l.log.GetLevel()

	if level <= Debug {
		elapsed := time.Since(begin)

		sql, rows := fc()
		if rows == -1 {
			l.log.Debug("%s\n[%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.log.Debug("%s\n[%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
