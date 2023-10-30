package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/service"

	"github.com/gin-gonic/gin"
)

var _ TransportLayer = (*transport)(nil)

type TransportLayer interface {
	Service() service.ServiceLayer
	healthCheck(c *gin.Context)

	// Auth
	signup(c *gin.Context)
	login(c *gin.Context)

	// Users
	createUser(c *gin.Context)
	getUser(c *gin.Context)
	updateUser(c *gin.Context)
	deleteUser(c *gin.Context)
	searchUsers(c *gin.Context)
	changePassword(c *gin.Context)

	// User Posts
	createUserPost(c *gin.Context)
}

func New(service service.ServiceLayer) *transport {
	return &transport{
		service: service,
	}
}

type transport struct {
	service service.ServiceLayer
}

func (t transport) Service() service.ServiceLayer {
	return t.service
}

func (t transport) healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "API is up and running :)"})
}
