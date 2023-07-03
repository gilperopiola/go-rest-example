package main

import (
	"log"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service"
	"github.com/gilperopiola/go-rest-example/pkg/transport"
)

func main() {

	// Create dependencies
	var config config.Config
	var database repository.Database
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
	service := service.Service{
		Repository: &repository,
		Codec:      &codec,
	}

	// Set up endpoints & router
	endpointsHandler := transport.Endpoints{
		Service:      &service,
		ErrorsMapper: &transport.ErrorsMapper{},
	}

	router.Setup(endpointsHandler)

	// Start server
	log.Println("server started")
	router.Run(":" + os.Getenv("PORT"))
}
