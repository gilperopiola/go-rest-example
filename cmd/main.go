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

	log.Println("Server started")

	// Setup dependencies
	var (
		config     = config.NewConfig()
		codec      = codec.NewCodec()
		database   = repository.NewDatabase(config.DATABASE)
		repository = repository.NewRepository(database)
		service    = service.NewService(&repository, &codec, config, service.ErrorsMapper{})
		endpoints  = transport.NewEndpoints(service, &codec, transport.ErrorsMapper{})
		router     = transport.NewRouter(endpoints, config)
	)

	defer database.Close()

	// Start server
	log.Println("Server running")
	router.Run(":" + os.Getenv("PORT"))
}
