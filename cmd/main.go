package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

const appName = "go-rest-example"

func main() {

	// It all starts here!
	log.Println("Server started")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.NewConfig()

		// Initialize logger
		logger = logrus.New()

		// Initialize monitoring (New Relic)
		monitoring = NewMonitoring(config)

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.SECRET, config.JWT.SESSION_DURATION_DAYS)

		// Setup codec for encoding and decoding
		codec = codec.NewCodec()

		// Establish database connection
		database = repository.NewDatabase(config.DATABASE, logger)

		// Initialize repository with the database connection
		repository = repository.NewRepository(database)

		// Setup the main service with dependencies
		service = service.NewService(repository, auth, codec, config, service.NewErrorsMapper())

		// Setup endpoints & transport layer with dependencies
		endpoints = transport.NewTransport(service, codec, transport.NewErrorsMapper(logger))

		// Initialize the router with the endpoints
		router = transport.NewRouter(endpoints, config, auth, logger, monitoring)
	)

	// Defer closing open connections
	defer database.Close()

	// Set log format and level
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)

	// Start server
	log.Println("Running server on port " + config.PORT)

	err := router.Run(":" + config.PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Have a nice day!
}

func NewMonitoring(config config.ConfigInterface) *newrelic.Application {
	if !config.GetMonitoringConfig().ENABLED {
		return nil
	}

	// If monitoring is enabled, use license to create New Relic app
	license := config.GetMonitoringConfig().SECRET
	if license == "" {
		log.Fatalf("New Relic license not found")
	}

	newRelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(license),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	// Panic on failure
	if err != nil {
		log.Fatalf("Failed to start New Relic: %v", err)
	}

	return newRelicApp
}
