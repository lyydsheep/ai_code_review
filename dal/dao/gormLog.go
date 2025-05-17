package dao

import (
	"context"
	log "github.com/faiz/llm-code-review/common/logger"
	"gorm.io/gorm/logger"
	"time"
)

var _ logger.Interface = (*gormLogger)(nil)
var _GormLogger *gormLogger

type gormLogger struct {
	SlowThreshold time.Duration
}

func InitGormLogger() {
	_GormLogger = &gormLogger{
		// 可以选择放入配置文件
		SlowThreshold: 100 * time.Millisecond,
	}
}

func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return _GormLogger
}

func (g *gormLogger) Info(c context.Context, msg string, data ...interface{}) {
	log.New(c).Info(msg, "data", data)
}

func (g *gormLogger) Warn(c context.Context, msg string, data ...interface{}) {
	log.New(c).Warn(msg, "data", data)
}

func (g *gormLogger) Error(c context.Context, msg string, data ...interface{}) {
	log.New(c).Error(msg, "data", data)
}

func (g *gormLogger) Debug(c context.Context, msg string, data ...interface{}) {
	log.New(c).Debug(msg, "data", data)
}

func (g *gormLogger) Trace(c context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	cost := time.Since(begin).Milliseconds()
	sql, rows := fc()
	if err != nil {
		g.Error(c, "error sql", "cost(ms)", cost, "sql", sql, "rows", rows, "err", err)
	}
	if cost > g.SlowThreshold.Milliseconds() {
		g.Warn(c, "slow sql", "cost(ms)", cost, "sql", sql, "rows", rows)
	} else {
		// 使用 debug 级别，减少实际生产环境中的日志数量
		g.Debug(c, "sql", "cost(ms)", cost, "sql", sql, "rows", rows)
	}
}
