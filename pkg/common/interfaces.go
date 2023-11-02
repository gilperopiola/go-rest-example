package common

import "github.com/sirupsen/logrus"

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
	Query(key string) (value string)
	DefaultQuery(key string, defaultValue string) string
}

type LoggerI interface {
	Print(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	WithField(key string, value interface{}) *logrus.Entry

	Fatalf(format string, args ...interface{})
}
