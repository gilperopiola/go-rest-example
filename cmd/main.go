package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/metrics"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
)

// TODO
// - NewRelic lo instancio por request
// - Redis
// - More tests
// - Add swagger docs
// - pprof -> check goroutine leaks
// - batch insert?
// - change password
// - reset password
// - roles to DB
// - prometheus

func main() {

	// It all starts here
	log.Println("Server starting ;)")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.New(".env")

		// Init Prometheus metrics
		metrics = metrics.New()

		// Initialize logger
		logger = middleware.NewLogger()

		// Initialize middlewares: logger, monitoring, prometheus
		middlewares = middleware.Middlewares{
			LoggerToCtx: middleware.NewLoggerToContextMiddleware(logger),
			Monitoring:  middleware.NewMonitoringMiddleware(config.Monitoring),
			Prometheus:  middleware.NewPrometheusMiddleware(metrics),
		}

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.Secret, config.JWT.SessionDurationDays)

		// Establish database connection
		database = repository.NewDatabase(*config, logger)

		// Initialize repository layer with the database connection
		repository = repository.New(database)

		// Setup the main service layer with dependencies
		service = service.New(repository, auth, config)

		// Setup endpoints & transport layer with dependencies
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
