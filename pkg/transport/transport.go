package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	Service      service.ServiceProvider
	Codec        codec.CodecProvider
	ErrorsMapper ErrorsMapperProvider
}

type TransportProvider interface {

	// Auth

	Signup(c *gin.Context)
	Login(c *gin.Context)

	// Users

	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

func NewTransport(service service.ServiceProvider, codec codec.CodecProvider, errorsMapper ErrorsMapperProvider) Transport {
	return Transport{
		Service:      service,
		Codec:        codec,
		ErrorsMapper: errorsMapper,
	}
}
