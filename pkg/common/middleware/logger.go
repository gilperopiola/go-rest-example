package middleware

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/sirupsen/logrus"
	gormLogger "gorm.io/gorm/logger"
)

type LoggerAdapter struct {
	Logger
}

type Logger struct {
	*logrus.Logger
}

func NewLogger(logInfo bool) *LoggerAdapter {

	level := logrus.ErrorLevel
	if logInfo {
		level = logrus.InfoLevel
	}

	return &LoggerAdapter{
		Logger: Logger{
			Logger: &logrus.Logger{
				Out: os.Stdout, Formatter: &CustomFormatter{},
				Hooks: make(logrus.LevelHooks), Level: level,
			},
		},
	}
}

func (l *Logger) Error(msg string, context map[string]interface{}) {
	l.prepareLogger(&msg, context).Error(msg)
}

func (l *Logger) Warn(msg string, context map[string]interface{}) {
	l.prepareLogger(&msg, context).Warn(msg)
}

func (l *Logger) Info(msg string, context map[string]interface{}) {
	l.prepareLogger(&msg, context).Info(msg)
}

func (l *Logger) Debug(msg string, context map[string]interface{}) {
	l.prepareLogger(&msg, context).Debug(msg)
}

func (l *Logger) DebugEnabled() bool {
	return l.IsLevelEnabled(logrus.DebugLevel)
}

// prepareLogger adds the necessary fields to the log and a new line to the message if it's not there
func (l *Logger) prepareLogger(msg *string, context map[string]interface{}) *logrus.Entry {

	// Add fields to log
	log := l.Logger.WithField("msg", *msg)
	for k, v := range context {
		log = log.WithField(k, v)
	}

	// New Relic
	if *msg == "application created" || *msg == "application connected" || *msg == "final configuration" ||
		*msg == "collector message" || *msg == "harvest failure" || *msg == "application connect failure" {
		log = log.WithField("from", common.NewRelic.String())
	}

	// Gin
	if strings.Contains(*msg, "[GIN-debug]") || strings.Contains(*msg, "GIN_MODE") || strings.Contains(*msg, "gin.SetMode") {
		log = log.WithField("from", common.Gin.String())
	}

	// Add new line if it's not there
	if !strings.Contains(*msg, "\n") { // TODO endswith
		*msg += "\n"
	}

	return log
}

/*----------------------------
//    GORM LOGGER ADAPTER
//
//	This adapter is used to unify our Logger with GORM's Logger
//--------------------------------------------------------------*/

func (l *LoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		l.Logger.Error(
			sql,
			map[string]interface{}{
				"from":    common.Gorm.String(),
				"rows":    rows,
				"elapsed": float64(elapsed.Nanoseconds()) / 1e6,
				"err":     err,
			},
		)
	} else {
		l.Logger.Info(
			sql,
			map[string]interface{}{
				"from":    common.Gorm.String(),
				"rows":    rows,
				"elapsed": float64(elapsed.Nanoseconds()) / 1e6,
			},
		)
	}
}

func (l *LoggerAdapter) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	switch level {
	case gormLogger.Info:
		l.Logger.SetLevel(logrus.InfoLevel)
	case gormLogger.Warn:
		l.Logger.SetLevel(logrus.WarnLevel)
	case gormLogger.Error:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case gormLogger.Silent:
		l.Logger.SetLevel(logrus.PanicLevel)
	}
	return l
}

func (l *LoggerAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Info(msg, map[string]interface{}{})
}

func (l *LoggerAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Warn(msg, map[string]interface{}{})
}

func (l *LoggerAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Logger.Error(msg, map[string]interface{}{})
}

func (l *LoggerAdapter) Write(p []byte) (n int, err error) {
	l.Logger.Info(string(p), map[string]interface{}{})
	return len(p), nil
}
