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
// - Batch insert
// - Reset password
// - Roles to DB
// - Prometheus enable flag
// - Request IDs
// - Logic from DeleteUser to service layer
// - Search & Fix TODOs
// - Replace user.Exists when you can
// - Change in config JWT -> Auth
// - OpenAPI (Swagger)

// Note: This is the entrypoint of the application.
// The HTTP Requests entrypoint is the Prometheus HandlerFunc in prometheus.go

func main() {

	log.Println("Server starting ;)")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.New()

		// Initialize logger
		logger = middleware.NewLogger()

		// Initialize middlewares
		middlewares = []gin.HandlerFunc{
			gin.Recovery(), // Panic recovery
			middleware.NewRateLimiterMiddleware(middleware.NewRateLimiter(200)),                     // Rate Limiter
			middleware.NewCORSConfigMiddleware(),                                                    // CORS
			middleware.NewNewRelicMiddleware(middleware.NewNewRelic(config.Monitoring, logger)),     // New Relic (monitoring)
			middleware.NewPrometheusMiddleware(middleware.NewPrometheus(config.Monitoring, logger)), // Prometheus (metrics)
			middleware.NewTimeoutMiddleware(config.General.Timeout),                                 // Timeout
			middleware.NewErrorHandlerMiddleware(logger),                                            // Error Handler
		}

		// Initialize authentication module
		auth = auth.New(config.Auth.JWTSecret, config.Auth.SessionDurationDays)

		// Establish database connection
		database = repository.NewDatabase(config, repository.NewDBLogger(logger.Out))

		// Initialize repositoryLayer layer with the database connection
		repositoryLayer = repository.New(database)

		// Setup the main serviceLayer layer
		serviceLayer = service.New(repositoryLayer, auth, config)

		// Setup endpoints & transport layer
		transportLayer = transport.New(serviceLayer)

		// Initialize the router with the endpoints
		router = transport.NewRouter(transportLayer, *config, auth, middlewares...)
	)

	// Start server
	port := config.General.Port
	log.Println("Running server on port " + port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Have a nice day!
}
