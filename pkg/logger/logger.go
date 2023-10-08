package logger

import "github.com/sirupsen/logrus"

type LoggerI interface {
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func NewLogger() LoggerI {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.ErrorLevel)
	return logger
}
