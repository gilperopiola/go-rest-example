package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

// Compile time check to validate that the transport struct implements the TransportLayer interface
var _ TransportLayer = (*transport)(nil)

type TransportLayer interface {
	signup(c *gin.Context)
	login(c *gin.Context)
	createUser(c *gin.Context)
	getUser(c *gin.Context)
	updateUser(c *gin.Context)
	deleteUser(c *gin.Context)
	searchUsers(c *gin.Context)
	changePassword(c *gin.Context)
	createUserPost(c *gin.Context)

	healthCheck(c *gin.Context)
}

type transport struct {
	service.ServiceLayer
}

func New(service service.ServiceLayer) *transport {
	return &transport{service}
}
