package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
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

type Request interface {
	entities.SignupRequest |
		entities.LoginRequest |
		entities.GetUserRequest |
		entities.UpdateUserRequest |
		entities.DeleteUserRequest
}
type Response interface {
	entities.SignupResponse |
		entities.LoginResponse |
		entities.GetUserResponse |
		entities.UpdateUserResponse |
		entities.DeleteUserResponse
}

// HandleRequest takes:
//
//   - a transport and a gin context
//   - a function that makes a request from the gin context
//   - a function that calls the service with that request
//
// It returns a response with the result of the service call.
func HandleRequest[req Request, resp Response](t Transport, c *gin.Context, makeRequest func(*gin.Context) (req, error), serviceCall func(req) (resp, error)) {

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
