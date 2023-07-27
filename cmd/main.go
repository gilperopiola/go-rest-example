package main

import (
	"log"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
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
		auth       = auth.NewAuth(config.JWT.SECRET, config.JWT.SESSION_DURATION_DAYS)
		codec      = codec.NewCodec()
		database   = repository.NewDatabase(config.DATABASE)
		repository = repository.NewRepository(database)
		service    = service.NewService(repository, auth, codec, config, service.ErrorsMapper{})
		endpoints  = transport.NewTransport(service, codec, transport.ErrorsMapper{})
		router     = transport.NewRouter(endpoints, config, auth)
	)

	defer database.Close()

	// Start server
	log.Println("Server running")
	router.Run(":" + os.Getenv("PORT"))
}
