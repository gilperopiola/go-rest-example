package main

import (
	"log"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/codec"
	cfg "github.com/gilperopiola/go-rest-example/pkg/config"
	repository "github.com/gilperopiola/go-rest-example/pkg/repository"
	service_v1 "github.com/gilperopiola/go-rest-example/pkg/service"
	transport "github.com/gilperopiola/go-rest-example/pkg/transport"
)

var config cfg.Config
var database repository.Database
var service service_v1.ServiceHandler
var router transport.Router

func main() {
	// Set up configuration
	config.Setup()

	// Set up database
	database.Setup(config.DATABASE)
	defer database.Close()

	// Set up repository
	repository := repository.RepositoryHandler{Database: &database}

	// Set up codec
	codec := codec.CodecHandler{}

	// Set up service
	service = service_v1.ServiceHandler{
		Database:   &database,
		Repository: &repository,
		Codec:      &codec,
	}

	// Set up endpoints & router
	endpointsHandler := transport.EndpointsHandler{
		Database: &database,
		Service:  &service,
	}
	router.Setup(endpointsHandler)

	// Start server
	log.Println("server started")
	router.Run(":" + os.Getenv("PORT"))
}
