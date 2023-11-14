package common

import (
	"context"
	"time"

	gormLogger "gorm.io/gorm/logger"
)

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}

type LoggerI interface {
	Error(ctx context.Context, msg string, data ...interface{})
	Info(ctx context.Context, msg string, data ...interface{})
	LogMode(level gormLogger.LogLevel) gormLogger.Interface
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
	Warn(ctx context.Context, msg string, data ...interface{})
	Write(p []byte) (n int, err error)
	Print(v ...interface{})
}
