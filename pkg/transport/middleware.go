package transport

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// This is exported because the monitoring is initialized in the main.go file
func NewMonitoringMiddleware(config config.ConfigI) gin.HandlerFunc {
	monitoringConfig := config.GetMonitoringConfig()

	// If monitoring is not enabled, return empty middleware
	if !monitoringConfig.ENABLED {
		return gin.HandlerFunc(func(c *gin.Context) {})
	}

	// If monitoring is enabled, use license to create New Relic app
	license := monitoringConfig.SECRET
	if license == "" {
		log.Fatalf("New Relic license not found")
	}

	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(monitoringConfig.APP_NAME),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	// Panic on failure
	if err != nil {
		log.Fatalf("Failed to start New Relic: %v", err)
	}

	return nrgin.Middleware(newRelicApp)
}

func newLoggerToContextMiddleware(logger logger.LoggerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func newCORSConfigMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
