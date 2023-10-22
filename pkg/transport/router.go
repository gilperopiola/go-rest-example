package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/middleware"

	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
}

func NewRouter(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares ...gin.HandlerFunc) router {
	var router router
	router.Setup(t, cfg, auth, middlewares...)
	return router
}

func (router *router) Setup(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares ...gin.HandlerFunc) {

	// Create router. Set debug/release mode
	router.prepare(!cfg.Debug)

	// Add middleware
	router.Use(gin.Recovery())                               // Panic recovery
	router.Use(middleware.NewTimeoutMiddleware(cfg.Timeout)) // Timeout
	router.Use(middleware.NewCORSConfigMiddleware())         // CORS
	for _, m := range middlewares {
		router.Use(m) // Monitoring & Adding Logger to Ctx
	}

	// Set endpoints
	router.setPublicEndpoints(t)
	router.setUserEndpoints(t, auth)
	router.setAdminEndpoints(t, auth)
}

func (router *router) prepare(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Engine = gin.New()
	router.Engine.SetTrustedProxies(nil)
}
