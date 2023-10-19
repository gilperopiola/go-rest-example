package middleware

import (
	"log"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type LoggerI interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
}

func NewLogger() LoggerI {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func NewMonitoringMiddleware(config config.Monitoring) gin.HandlerFunc {
	// If monitoring is not enabled, return empty middleware
	if !config.Enabled {
		return gin.HandlerFunc(func(c *gin.Context) {})
	}

	// If monitoring is enabled, use license to create New Relic app
	license := config.Secret
	if license == "" {
		log.Fatalf("New Relic license not found")
	}

	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.AppName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	// Panic on failure
	if err != nil {
		log.Fatalf("Failed to start New Relic: %v", err)
	}

	return nrgin.Middleware(newRelicApp)
}

func NewTimeoutMiddleware(timeoutSeconds int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(timeoutSeconds)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
	)
}

func NewLoggerToContextMiddleware(logger LoggerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func NewCORSConfigMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
