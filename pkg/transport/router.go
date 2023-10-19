package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(t TransportLayer, cfg config.General, auth auth.AuthI, logger middleware.LoggerI, monitoring gin.HandlerFunc) Router {
	var router Router
	router.Setup(t, cfg, auth, logger, monitoring)
	return router
}

func (router *Router) Setup(t TransportLayer, cfg config.General, auth auth.AuthI, logger middleware.LoggerI, monitoring gin.HandlerFunc) {

	// Create router. Set debug/release mode
	router.prepare(!cfg.Debug)

	// Add middleware
	router.Use(monitoring)                                      // Monitoring
	router.Use(gin.Recovery())                                  // Panic recovery
	router.Use(middleware.NewTimeoutMiddleware(cfg.Timeout))    // Timeout
	router.Use(middleware.NewCORSConfigMiddleware())            // CORS
	router.Use(middleware.NewLoggerToContextMiddleware(logger)) // Logger to context

	// Set endpoints
	router.setPublicEndpoints(t)
	router.setUserEndpoints(t, auth)
	router.setAdminEndpoints(t, auth)
}

func (router *Router) prepare(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Engine = gin.New()
	router.Engine.SetTrustedProxies(nil)
}
