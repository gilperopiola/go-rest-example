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

/* ----------------- */

type RequestProvider interface {
	Validate() error
}

type Request RequestProvider
type Response interface{}

func HandleRequest[T Request, R Response](t Transport, c *gin.Context, makeRequest func(*gin.Context) (T, error), serviceCall func(T) (R, error)) {

	// Make, validate and get request
	request, err := makeRequest(c)
	if err != nil {
		c.JSON(t.ErrorsMapper.Map(err))
		return
	}

	// Call service with that request
	response, err := serviceCall(request)
	if err != nil {
		c.JSON(t.ErrorsMapper.Map(err))
		return
	}

	// Return OK
	c.JSON(returnOK(response))
}
