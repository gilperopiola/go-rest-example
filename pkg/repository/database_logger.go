package repository

import (
	"context"
	"io"
	"log"
	"time"

	c "github.com/gilperopiola/go-rest-example/pkg/common"

	"gorm.io/gorm/logger"
)

// Logrus logger is not compatible with GORM v2
type DBLogger struct {
	*log.Logger
}

func NewDBLogger(out io.Writer) *DBLogger {
	prefix := c.WhiteBold + "\r\n"
	logger := log.New(out, prefix, log.LstdFlags)
	return &DBLogger{logger}
}

// Trace is a callback that can be registered with GORM to track every SQL query
func (l *DBLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		layout := "%s[ERR]%s %s%s%s (%d rows affected) %s[%.3fms] %s%v%s\n"
		l.Printf(layout, c.RedBold, c.Reset, c.White, sql, c.Reset, rows, c.Green, float64(elapsed.Nanoseconds())/1e6, c.Red, err, c.Reset)
	} else {
		layout := "%s[SQL]%s %s%s%s (%d rows affected) %s[%.3fms]%s\n"
		l.Printf(layout, c.YellowBold, c.Reset, c.White, sql, c.Gray, rows, c.Green, float64(elapsed.Nanoseconds())/1e6, c.Reset)
	}
}

func (l *DBLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	return &newLogger
}

func (l *DBLogger) Info(ctx context.Context, msg string, data ...interface{}) {}

func (l *DBLogger) Warn(ctx context.Context, msg string, data ...interface{}) {}

func (l *DBLogger) Error(ctx context.Context, msg string, data ...interface{}) {}
