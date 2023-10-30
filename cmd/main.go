package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
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
// - reset password
// - roles to DB
// - Fix Readme
// - Prometheus enable flag
// - Fix errors handling so that we can have errors with parameters
// - Request IDs
// - Requests pkg divide by endpoint files

// Note: This is the entrypoint of the application.
// The HTTP Requests entrypoint is the Prometheus HandlerFunc in prometheus.go
func main() {

	// It all starts here
	log.Println("Server starting ;)")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.New()

		// Initialize logger
		logger = common.NewLogger()

		// Initialize middlewares
		middlewares = []gin.HandlerFunc{
			gin.Recovery(), // Panic recovery
			middleware.NewRateLimiterMiddleware(middleware.NewRateLimiter(200)),         // Rate Limiter
			middleware.NewCORSConfigMiddleware(),                                        // CORS
			middleware.NewNewRelicMiddleware(middleware.NewNewRelic(config.Monitoring)), // New Relic (monitoring)
			middleware.NewPrometheusMiddleware(middleware.NewPrometheus()),              // Prometheus (metrics)
			middleware.NewTimeoutMiddleware(config.General.Timeout),                     // Timeout
			middleware.NewErrorHandlerMiddleware(logger),                                // Error Handler
		}

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.Secret, config.JWT.SessionDurationDays)

		// Establish database connection
		database = repository.NewDatabase(config)

		// Initialize repository layer with the database connection
		repository = repository.New(database)

		// Setup the main service layer
		service = service.New(repository, auth, config)

		// Setup endpoints & transport layer
		endpoints = transport.New(service)

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
