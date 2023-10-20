package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type TransportLayer interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)

	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	CreateUserPost(c *gin.Context)
}

type Transport struct {
	Service      service.ServiceLayer
	ErrorsMapper errorsMapperI
}

func New(service service.ServiceLayer, errorsMapper errorsMapperI) Transport {
	return Transport{
		Service:      service,
		ErrorsMapper: errorsMapper,
	}
}
