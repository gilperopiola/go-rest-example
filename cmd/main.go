package main

import (
	"log"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/codec"
	cfg "github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	service_v1 "github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
)

func main() {

	// Create dependencies
	var config cfg.Config
	var database repository.Database
	var service service_v1.Service
	var router transport.Router

	// Set up configuration
	config.Setup()

	// Set up database
	database.Setup(config.DATABASE)
	defer database.Close()

	// Set up repository
	repository := repository.Repository{
		Database: database,
	}

	// Set up codec
	codec := codec.Codec{}

	// Set up service
	service = service_v1.Service{
		Database:   &database,
		Repository: &repository,
		Codec:      &codec,
	}

	// Set up endpoints & router
	endpointsHandler := transport.Endpoints{
		Database:     &database,
		Service:      &service,
		ErrorsMapper: &transport.ErrorsMapper{},
	}

	router.Setup(endpointsHandler)

	// Start server
	log.Println("server started")
	router.Run(":" + os.Getenv("PORT"))
}
