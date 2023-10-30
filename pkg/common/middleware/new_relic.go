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

func NewNewRelic(config config.Monitoring) *newrelic.Application {

	// If monitoring is not enabled, return empty app
	if !config.Enabled {
		return nil
	}

	// If monitoring is enabled, use license to create New Relic app
	license := config.Secret
	if license == "" {
		log.Fatalf("New Relic license not found")
	}

	// Create app
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.AppName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	// Panic on failure
	if err != nil {
		log.Fatalf("Failed to start New Relic: %v", err)
	}

	return newRelicApp
}
