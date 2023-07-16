package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	Service      service.ServiceIFace
	Codec        codec.CodecIFace
	ErrorsMapper ErrorsMapperIface
}

type EndpointsIface interface {

	// Auth

	Signup(c *gin.Context)
	Login(c *gin.Context)

	// Users

	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

func NewEndpoints(service service.ServiceIFace, codec codec.CodecIFace, errorsMapper ErrorsMapperIface) Endpoints {
	return Endpoints{
		Service:      service,
		Codec:        codec,
		ErrorsMapper: errorsMapper,
	}
}
