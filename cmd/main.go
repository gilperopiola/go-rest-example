package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
	"github.com/gin-gonic/gin"
)

// TODO
// - Redis
// - More tests
// - Add swagger docs
// - batch insert?
// - change password
// - reset password
// - roles to DB
// - Fix Readme

// Note: The HTTP Requests entrypoint is the Prometheus HandlerFunc

func main() {

	// It all starts here
	log.Println("Server starting ;)")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.New(".env")

		// Initialize logger
		logger = middleware.NewLogger()

		// We use prometheus to get metrics
		prometheus = middleware.NewPrometheus(logger)

		// We use New Relic to monitor the app
		newRelic = middleware.NewMonitoringNewRelic(config.Monitoring)

		// Initialize middlewares
		middlewares = []gin.HandlerFunc{
			gin.Recovery(),
			middleware.NewCORSConfigMiddleware(),
			middleware.NewNewRelicMiddleware(newRelic),
			middleware.NewPrometheusMiddleware(prometheus),
			middleware.NewTimeoutMiddleware(config.General.Timeout),
		}

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.Secret, config.JWT.SessionDurationDays)

		// Establish database connection
		database = repository.NewDatabase(*config, logger)

		// Initialize repository layer with the database connection
		repository = repository.New(database)

		// Setup the main service layer
		service = service.New(repository, auth, config)

		// Setup endpoints & transport layer
		endpoints = transport.New(service, transport.NewErrorsMapper(logger))

		// Initialize the router with the endpoints
		router = transport.NewRouter(endpoints, config.General, auth, middlewares)
	)

	// Defer closing open connections
	defer database.Close()

	// Start server
	port := config.General.Port
	log.Println("Running server on port " + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Have a nice day!
}
