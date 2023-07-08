package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

type Endpoints struct {
	Service      service.ServiceIFace
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
