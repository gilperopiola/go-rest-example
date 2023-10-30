package common

import "github.com/sirupsen/logrus"

type LoggerI interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	WithField(key string, value interface{}) *logrus.Entry

	Fatalf(format string, args ...interface{})
}

func NewLogger() LoggerI {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}
