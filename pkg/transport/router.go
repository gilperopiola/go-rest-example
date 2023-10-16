package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/logger"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(transport TransportLayer, cfg config.GeneralConfig, auth auth.AuthI, logger logger.LoggerI, monitoring gin.HandlerFunc) Router {
	var router Router
	router.Setup(transport, cfg, auth, logger, monitoring)
	return router
}

func (router *Router) Setup(transport TransportLayer, cfg config.GeneralConfig, auth auth.AuthI, logger logger.LoggerI, monitoring gin.HandlerFunc) {

	// Create router. Set debug/release mode
	router.Prepare(!cfg.Debug)

	// Add middleware
	router.Use(monitoring)                                      // Monitoring
	router.Use(gin.Recovery())                                  // Panic recovery
	router.Use(middleware.NewTimeoutMiddleware(cfg.Timeout))    // Timeout
	router.Use(middleware.NewCORSConfigMiddleware())            // CORS
	router.Use(middleware.NewLoggerToContextMiddleware(logger)) // Logger to context

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
