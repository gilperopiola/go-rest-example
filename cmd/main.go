package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/logger"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
)

func main() {

	// It all starts here
	log.Println("Server started")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.NewConfig()

		// Initialize logger
		logger = logger.NewLogger()

		// Initialize monitoring as middleware (New Relic)
		monitoringMiddleware = transport.NewMonitoringMiddleware(config)

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.SECRET, config.JWT.SESSION_DURATION_DAYS)

		// Establish database connection
		database = repository.NewDatabase(config.DATABASE, logger)

		// Initialize repository layer with the database connection
		repository = repository.NewRepository(database)

		// Setup the main service layer with dependencies
		service = service.NewService(repository, auth, config)

		// Setup endpoints & transport layer with dependencies
		endpoints = transport.NewTransport(service, transport.NewErrorsMapper(logger))

		// Initialize the router with the endpoints
		router = transport.NewRouter(endpoints, config, auth, logger, monitoringMiddleware)
	)

	// Defer closing open connections
	defer database.Close()

	// Start server
	log.Println("Running server on port " + config.PORT)

	err := router.Run(":" + config.PORT)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Have a nice day!
}
