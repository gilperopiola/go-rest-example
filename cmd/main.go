package main

import (
	"context"
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"
	"github.com/gilperopiola/go-rest-example/pkg/service"

	//repository "github.com/gilperopiola/go-rest-example/pkg/sql_repository"
	mongoRepository "github.com/gilperopiola/go-rest-example/pkg/mongo_repository"
	"github.com/gilperopiola/go-rest-example/pkg/transport"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/*---------------------------------------------
// This is the entrypoint of the application.
//
// The HTTP Requests entrypoint is the RateLimiter middleware in /middleware/rate_limiter.go
/-------------------------------------------------------------------------------------------*/

func main() {

	log.Println("Server starting ;)")

	/*-------------------------
	//      Dependencies
	//-----------------------*/

	config := config.New()
	common.SetConfig(config)
	log.Println("Config OK!")

	logger := middleware.NewLogger(config.LogInfo)
	common.SetLogger(logger)
	logger.Logger.Info("Logger OK!", nil)

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

	//database := repository.NewDatabase()
	//sqlDatabase := database.SQLDB()
	mongoDatabase := mongoRepository.NewDatabase()
	defer func() {
		if err := mongoDatabase.DB().Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	logger.Logger.Info("Database OK!", nil)

	//repositoryLayer := repository.New(database)
	logger.Logger.Info("Repository Layer OK!", nil)

	mongoRepositoryLayer := mongoRepository.New(mongoDatabase)
	logger.Logger.Info("Mongo Repository Layer OK!", nil)

	serviceLayer := service.New(mongoRepositoryLayer)
	logger.Logger.Info("Service Layer OK!", nil)

	transportLayer := transport.New(serviceLayer, validator.New(), nil, mongoDatabase.DB())
	logger.Logger.Info("Transport Layer OK!", nil)

	router := transport.NewRouter(transportLayer, middlewares...)
	logger.Logger.Info("Router & Endpoints OK!", nil)

	/*---------------------------
	//       Server Start
	//-------------------------*/

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
