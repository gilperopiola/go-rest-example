package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

var _ TransportLayer = (*transport)(nil)

type TransportLayer interface {
	Service() service.ServiceLayer
	ErrorsMapper() errorsMapperI

	healthCheck(c *gin.Context)

	signup(c *gin.Context)
	login(c *gin.Context)

	createUser(c *gin.Context)
	getUser(c *gin.Context)
	updateUser(c *gin.Context)
	deleteUser(c *gin.Context)
	searchUsers(c *gin.Context)

	createUserPost(c *gin.Context)
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
