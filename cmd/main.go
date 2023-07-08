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

	// Config, Database, Repository, Codec
	config.Setup()
	database.Setup(config.DATABASE)
	defer database.Close()
	repository := repository.Repository{Database: database}
	codec := codec.Codec{}

	// Service, Endpoints, Router
	service := service.NewService(&repository, &codec, config, service.ErrorsMapper{})
	endpointsHandler := transport.Endpoints{Service: service, ErrorsMapper: &transport.ErrorsMapper{}}
	router.Setup(endpointsHandler, config.JWT)

	// Start server
	log.Println("server started")
	router.Run(":" + os.Getenv("PORT"))
}
