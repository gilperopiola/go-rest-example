package middleware

import (
	"fmt"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewNewRelicMiddleware(app *newrelic.Application) gin.HandlerFunc {
	return nrgin.Middleware(app)
}

func NewNewRelic(config config.Monitoring, logger *LoggerAdapter) *newrelic.Application {

	// If New Relic is not enabled, return empty app
	if !config.NewRelicEnabled {
		logger.Logger.Info("New Relic Disabled", map[string]interface{}{"from": common.NewRelic.String()})
		return nil
	}

	// If monitoring is enabled, use license to create New Relic app
	license := config.NewRelicLicenseKey
	if license == "" {
		logger.Logger.Error("New Relic license not found", map[string]interface{}{"from": common.NewRelic.String()})
		os.Exit(1)
	}

	// Create app
	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicAppName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigLogger(&logger.Logger),
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	// Panic on failure
	if err != nil {
		logger.Logger.Info(fmt.Sprintf("Failed to start New Relic: %v", err), map[string]interface{}{"from": common.NewRelic.String()})
		os.Exit(1)
	}

	return newRelicApp
}
