package main

import (
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
)

// TODO
// - Redis
// - More tests
// - Add postman collection
// - Add swagger docs
// - pprof -> check goroutine leaks
// - Add another model (user 1-n)
// - batch insert?
// - change password
// - reset password
// - roles to DB
// - List users ep
// - api versioning
// - prometheus

func main() {

	// It all starts here
	log.Println("Server starting ;)")

	// Setup dependencies
	var (
		// Load configuration settings
		config = config.New()

		// Initialize logger & logger middleware
		logger           = middleware.NewLogger()
		loggerMiddleware = middleware.NewLoggerToContextMiddleware(logger)

		// Initialize monitoring as middleware (New Relic)
		monitoringMiddleware = middleware.NewMonitoringMiddleware(config.Monitoring)

		// Initialize authentication module
		auth = auth.NewAuth(config.JWT.Secret, config.JWT.SessionDurationDays)

		// Establish database connection
		database = repository.NewDatabase(config.Database, logger)

		// Initialize repository layer with the database connection
		repository = repository.New(database)

		// Setup the main service layer with dependencies
		service = service.New(repository, auth, config)

		// Setup endpoints & transport layer with dependencies
		endpoints = transport.New(service, transport.NewErrorsMapper(logger))

		// Initialize the router with the endpoints
		router = transport.NewRouter(endpoints, config.General, auth, loggerMiddleware, monitoringMiddleware)
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
