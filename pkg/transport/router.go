package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Router struct {
	*gin.Engine
}

func NewRouter(transport TransportLayer, config config.ConfigInterface, auth auth.AuthInterface, logger *logrus.Logger,
	monitoring *newrelic.Application) Router {
	var router Router
	router.Setup(transport, config.GetDebugMode(), auth, logger, monitoring)
	return router
}

func (router *Router) Setup(transport TransportLayer, debugMode bool, auth auth.AuthInterface, logger *logrus.Logger,
	monitoring *newrelic.Application) {

	// Create router. Set debug/release mode
	router.Prepare(!debugMode)

	// Add middleware
	router.Use(nrgin.Middleware(monitoring))
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
	router.Engine.SetTrustedProxies(nil)
}
