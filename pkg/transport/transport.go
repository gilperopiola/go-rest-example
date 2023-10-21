package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

var _ TransportLayer = (*transport)(nil)

type TransportLayer interface {
	Service() service.ServiceLayer
	ErrorsMapper() errorsMapperI

	Signup(c *gin.Context)
	Login(c *gin.Context)

	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	SearchUsers(c *gin.Context)

	CreateUserPost(c *gin.Context)
}

type transport struct {
	service      service.ServiceLayer
	errorsMapper errorsMapperI
}

func New(service service.ServiceLayer, errorsMapper errorsMapperI) transport {
	return transport{
		service:      service,
		errorsMapper: errorsMapper,
	}
}

func (t transport) Service() service.ServiceLayer {
	return t.service
}

func (t transport) ErrorsMapper() errorsMapperI {
	return t.errorsMapper
}
