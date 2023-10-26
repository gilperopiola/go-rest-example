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

func NewRouter(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares middleware.Middlewares) router {
	var router router
	router.Setup(t, cfg, auth, middlewares)
	return router
}

func (router *router) Setup(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares middleware.Middlewares) {

	// Create router. Set debug/release mode
	router.prepare(!cfg.Debug)

	// Add middleware
	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	// Set endpoints
	router.setEndpoints(t, cfg, auth)
}

func (router *router) prepare(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}
	router.Engine = gin.New()
	router.Engine.SetTrustedProxies(nil)
}
