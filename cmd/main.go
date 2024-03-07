package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	mongoRepository "github.com/gilperopiola/go-rest-example/pkg/repository/mongo_repository"
	sqlRepository "github.com/gilperopiola/go-rest-example/pkg/repository/sql_repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
	This is the entrypoint of the application.

	The HTTP Requests entrypoint is the RateLimiter middleware in /middleware/rate_limiter.go
*/

type dependencies struct {
	config      *config.Config
	logger      *middleware.LoggerAdapter
	middlewares []gin.HandlerFunc

	mySQLDatabase       *sqlRepository.Database
	mySQLDatabaseObject *sql.DB
	mongoDatabase       *mongoRepository.Database

	repositoryLayer repository.RepositoryLayer
	serviceLayer    service.ServiceLayer
	transportLayer  transport.TransportLayer

	router transport.Router
}

func main() {

	log.Println("Server starting ;)")

	/*-------------------------------------------
	                Dependencies
	/*------------------------------------------*/

	d := &dependencies{}

	/* Config & Logger
	/*----------------*/

	d.config = config.New()
	d.logger = middleware.NewLogger(d.config.LogInfo)

	common.SetConfig(d.config)
	common.SetLogger(d.logger)

	log := func(msg string) { d.logger.Logger.Info(msg, nil) }
	log("Config & Logger OK!")

	/* Middlewares
	/*------------*/

	d.middlewares = []gin.HandlerFunc{
		gin.Recovery(), // Panic recovery
		//middleware.NewTimeoutMiddleware(d.config.Timeout),                               		 			 // Timeout T0D0 Fix 500
		middleware.NewRateLimiterMiddleware(middleware.NewRateLimiter(200)),                         // Rate Limiter
		middleware.NewCORSConfigMiddleware(),                                                        // CORS
		middleware.NewNewRelicMiddleware(middleware.NewNewRelic(d.config.Monitoring, d.logger)),     // New Relic (monitoring)
		middleware.NewPrometheusMiddleware(middleware.NewPrometheus(d.config.Monitoring, d.logger)), // Prometheus (metrics)
		middleware.NewErrorHandlerMiddleware(d.logger),                                              // Error Handler
	}
	log("Middlewares OK!")

	/* Database & Repository
	/*----------------------*/

	switch d.config.Database.Type {
	case "mysql":
		d.mySQLDatabase = sqlRepository.NewDatabase()
		d.mySQLDatabaseObject = d.mySQLDatabase.GetSQLDB()
		d.repositoryLayer = sqlRepository.New(d.mySQLDatabase)
	case "mongodb":
		d.mongoDatabase = mongoRepository.NewDatabase()
		defer d.mongoDatabase.Disconnect(context.Background())
		d.repositoryLayer = mongoRepository.New(d.mongoDatabase, d.config.Database.Mongo)
	default:
		d.logger.Logger.Fatalf("Invalid Database type: %s", d.config.Database.Type)
	}
	log("Database & Repository Layer OK!")

	/* Service & Transport
	/*--------------------*/

	d.serviceLayer = service.New(d.repositoryLayer)
	log("Service Layer OK!")

	d.transportLayer = transport.New(d.serviceLayer, validator.New(), d.mySQLDatabaseObject, d.mongoDatabase.DB())
	log("Transport Layer OK!")

	/* Router
	/*--------*/

	d.router = transport.NewRouter(d.transportLayer, d.middlewares...)
	log("Router & Endpoints OK!")

	/*--------------------------------
	            Server Start
	/*-------------------------------*/

	if err := d.router.Run(":" + d.config.Port); err != nil {
		d.logger.Logger.Fatalf("Failed to start Server: %v :(", err)
	}

	// Have a great day! :)
}

/*------
// T0D0
// - Redis
// - More tests
// - Batch insert
// - Reset password
// - Roles to DB
// - Request IDs
// - Search & Fix T0D0s
// - OpenAPI (Swagger) */
