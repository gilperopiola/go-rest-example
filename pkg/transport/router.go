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

func NewRouter(transport TransportLayer, config config.ConfigInterface, auth auth.AuthInterface, logger *logrus.Logger) Router {
	var router Router
	router.Setup(transport, config.GetDebugMode(), auth, logger)
	return router
}

func (router *Router) Setup(transport TransportLayer, debugMode bool, auth auth.AuthInterface, logger *logrus.Logger) {

	// Create router. Set debug/release mode
	router.Prepare(!debugMode)

	// Add middleware
	router.Use(gin.Recovery())
	router.Use(addCorsConfigMiddleware())
	router.Use(addLoggerToContextMiddleware(logger))

	// Set endpoints
	router.SetPublicEndpoints(transport)
	router.SetUserEndpoints(transport, auth)
	router.SetAdminEndpoints(transport, auth)
}

func (router *Router) Prepare(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Engine = gin.New()
}
