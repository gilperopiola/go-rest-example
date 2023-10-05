package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	*gin.Engine
}

func NewRouter(transport TransportProvider, config config.ConfigInterface, auth auth.AuthInterface,
	logger *logrus.Logger) Router {
	var router Router
	router.Setup(transport, config, auth, logger)
	return router
}

func (router *Router) Setup(transport TransportProvider, config config.ConfigInterface, auth auth.AuthInterface,
	logger *logrus.Logger) {

	// Prepare router. Mode should be set first
	if !config.GetDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Engine = gin.New()

	// Add middleware
	router.Use(addCorsConfigMiddleware())
	router.Use(addLoggerToContextMiddleware(logger))

	// Set endpoints
	router.SetPublicEndpoints(transport)
	router.SetUserEndpoints(transport, auth)
	router.SetAdminEndpoints(transport, auth)
}

func (router *Router) SetPublicEndpoints(transport TransportProvider) {
	public := router.Group("/")
	public.POST("/signup", transport.Signup)
	public.POST("/login", transport.Login)
}

func (router *Router) SetUserEndpoints(transport TransportProvider, auth auth.AuthInterface) {
	users := router.Group("/users", auth.ValidateToken())
	users.GET("/:user_id", transport.GetUser)
	users.PATCH("/:user_id", transport.UpdateUser)
	users.DELETE("/:user_id", transport.DeleteUser)
}

func (router *Router) SetAdminEndpoints(transport TransportProvider, auth auth.AuthInterface) {
	admin := router.Group("/admin", auth.ValidateRole(auth.GetAdminRole()))
	admin.POST("/user", transport.CreateUser)
}
