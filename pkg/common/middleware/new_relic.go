package middleware

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewNewRelicMiddleware(app *newrelic.Application) gin.HandlerFunc {
	return nrgin.Middleware(app)
}

func NewNewRelic(config config.Monitoring, logger *Logger) *newrelic.Application {

	// If New Relic is not enabled, return empty app
	if !config.NewRelicEnabled {
		log.Println("New Relic Disabled")
		return nil
	}

	// If monitoring is enabled, use license to create New Relic app
	license := config.NewRelicLicenseKey
	if license == "" {
		log.Fatalf("New Relic license not found")
	}

	// Create app
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicAppName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigLogger(logger),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	// Panic on failure
	if err != nil {
		log.Fatalf("Failed to start New Relic: %v", err)
	}

	return newRelicApp
}
