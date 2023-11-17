package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	mongoRepository "github.com/gilperopiola/go-rest-example/pkg/mongo_repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	repository "github.com/gilperopiola/go-rest-example/pkg/sql_repository"
	"github.com/gilperopiola/go-rest-example/pkg/transport"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*
	This is the entrypoint of the application.

	The HTTP Requests entrypoint is the RateLimiter middleware in /middleware/rate_limiter.go
*/

func main() {

	log.Println("Server starting ;)")

	/*-------------------------------------------
	||              Dependencies
	/*------------------------------------------*/

	/* Config & Logger
	/*----------------*/

	config := config.New()
	common.SetConfig(config)
	log.Println("Config OK!")

	logger := middleware.NewLogger(config.LogInfo)
	common.SetLogger(logger)
	logger.Logger.Info("Logger OK!", nil)

	/* Middlewares
	/*------------*/

	middlewares := []gin.HandlerFunc{
		gin.Recovery(), // Panic recovery
		//middleware.NewTimeoutMiddleware(config.Timeout),                               		 // Timeout TODO Fix 500
		middleware.NewRateLimiterMiddleware(middleware.NewRateLimiter(200)),                     // Rate Limiter
		middleware.NewCORSConfigMiddleware(),                                                    // CORS
		middleware.NewNewRelicMiddleware(middleware.NewNewRelic(config.Monitoring, logger)),     // New Relic (monitoring)
		middleware.NewPrometheusMiddleware(middleware.NewPrometheus(config.Monitoring, logger)), // Prometheus (metrics)
		middleware.NewErrorHandlerMiddleware(logger),                                            // Error Handler
	}
	logger.Logger.Info("Middlewares OK!", nil)

	/* Database & Repository
	/*----------------------*/

	var repositoryLayer repository.RepositoryLayer
	var mySQLDatabase *repository.Database
	var mySQLDatabaseObject *sql.DB
	var mongoDatabase *mongoRepository.Database

	switch config.Database.Type {
	case "mysql":
		mySQLDatabase = repository.NewDatabase()
		mySQLDatabaseObject = mySQLDatabase.GetSQLDB()
		repositoryLayer = repository.New(mySQLDatabase)
	case "mongodb":
		mongoDatabase := mongoRepository.NewDatabase()
		defer mongoDatabase.Disconnect(context.Background())
		repositoryLayer = mongoRepository.New(mongoDatabase, config.Database.Mongo)
	default:
		logger.Logger.Fatalf("Invalid database type: %s", config.Database.Type)
	}
	logger.Logger.Info("Database & Repository Layer OK!", nil)

	/* Service
	/*---------*/

	serviceLayer := service.New(repositoryLayer)
	logger.Logger.Info("Service Layer OK!", nil)

	/* Transport
	/*----------*/

	transportLayer := transport.New(serviceLayer, validator.New(), mySQLDatabaseObject, mongoDatabase.DB())
	logger.Logger.Info("Transport Layer OK!", nil)

	/* Router
	/*--------*/

	router := transport.NewRouter(transportLayer, middlewares...)
	logger.Logger.Info("Router & Endpoints OK!", nil)

	/*--------------------------------
	||          Server Start
	/*-------------------------------*/

	port := config.Port
	logger.Logger.Infof("Running Server on port %s!\n", port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Failed to start Server: %v :(", err)
	}

	// Have a great day! :)
}

/*------
// TODO
// - Redis
// - More tests
// - Batch insert
// - Reset password
// - Roles to DB
// - Request IDs
// - Search & Fix TODOs
// - OpenAPI (Swagger) */
