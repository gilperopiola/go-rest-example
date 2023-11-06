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

// Note: This is the entrypoint of the application.
// The HTTP Requests entrypoint is the Prometheus HandlerFunc in pkg/common/middleware/prometheus.go

func main() {

	log.Println("Server starting ;)")

	/*-------------------------
	//      DEPENDENCIES
	//------------------------*/

	config := config.New()
	log.Println("Config OK")

	logger := middleware.NewLogger(config.General.LogInfo)
	logger.Logger.Info("Logger OK", nil)

	middlewares := []gin.HandlerFunc{
		gin.Recovery(), // Panic recovery
		middleware.NewRateLimiterMiddleware(middleware.NewRateLimiter(200)),                     // Rate Limiter
		middleware.NewCORSConfigMiddleware(),                                                    // CORS
		middleware.NewNewRelicMiddleware(middleware.NewNewRelic(config.Monitoring, logger)),     // New Relic (monitoring)
		middleware.NewPrometheusMiddleware(middleware.NewPrometheus(config.Monitoring, logger)), // Prometheus (metrics)
		middleware.NewTimeoutMiddleware(config.General.Timeout),                                 // Timeout
		middleware.NewErrorHandlerMiddleware(logger),                                            // Error Handler
	}
	logger.Logger.Info("Middlewares OK", nil)

	auth := auth.New(config.Auth.JWTSecret, config.Auth.SessionDurationDays)
	logger.Logger.Info("Auth OK", nil)

	database := repository.NewDatabase(config, logger)
	logger.Logger.Info("Database OK", nil)

	repositoryLayer := repository.New(database)
	logger.Logger.Info("Repository Layer OK", nil)

	serviceLayer := service.New(repositoryLayer, auth, config)
	logger.Logger.Info("Service Layer OK", nil)

	transportLayer := transport.New(serviceLayer)
	logger.Logger.Info("Transport Layer OK", nil)

	router := transport.NewRouter(transportLayer, config, auth, logger, middlewares...)
	logger.Logger.Info("Router & Endpoints OK", nil)

	/*---------------------------
	//       START SERVER
	//--------------------------*/

	port := config.General.Port
	logger.Logger.Info("Running server on port "+port, nil)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	/* Have a great day! :) */
}

// TODO
// - Redis
// - More tests
// - Batch insert
// - Reset password
// - Roles to DB
// - Request IDs
// - Logic from DeleteUser to service layer
// - Search & Fix TODOs
// - Replace user.Exists when you can
// - OpenAPI (Swagger)
