package transport

import (
	"net/http/pprof"

	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

type router struct {
	*gin.Engine
}

func NewRouter(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares []gin.HandlerFunc) router {
	var router router
	router.setup(t, cfg, auth, middlewares)
	return router
}

func (router *router) setup(t TransportLayer, cfg config.General, auth auth.AuthI, middlewares []gin.HandlerFunc) {

	// Create router. Set debug/release mode
	router.prepare(!cfg.Debug)

	// Add middlewares
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

//-----------------------------
//     ROUTES / ENDPOINTS
//-----------------------------

func (router *router) setEndpoints(transport TransportLayer, cfg config.General, authI auth.AuthI) {

	// Standard endpoints
	router.GET("/health", transport.healthCheck)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// V1
	v1 := router.Group("/v1")
	{
		router.setV1Endpoints(v1, transport, authI)
	}

	// Profiling
	if cfg.Profiling {
		router.profiling()
	}
}

func (router *router) setV1Endpoints(v1 *gin.RouterGroup, transport TransportLayer, authI auth.AuthI) {

	// Auth
	v1.POST("/signup", transport.signup)
	v1.POST("/login", transport.login)

	// Users
	users := v1.Group("/users", authI.ValidateToken(auth.AnyRole, true))
	{
		users.GET("/:user_id", transport.getUser)
		users.PATCH("/:user_id", transport.updateUser)
		users.DELETE("/:user_id", transport.deleteUser)
		users.PATCH("/:user_id/password", transport.changePassword)

		// User posts
		posts := users.Group("/:user_id/posts")
		{
			posts.POST("", transport.createUserPost)
		}
	}

	// Admins
	admin := v1.Group("/admin", authI.ValidateToken(auth.AdminRole, false))
	{
		admin.POST("/user", transport.createUser)
		admin.GET("/users", transport.searchUsers)
	}
}

func (r *router) profiling() {
	pprofGroup := r.Group("/debug/pprof")
	pprofGroup.GET("/", gin.WrapF(pprof.Index))
	pprofGroup.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	pprofGroup.GET("/profile", gin.WrapF(pprof.Profile))
	pprofGroup.POST("/symbol", gin.WrapF(pprof.Symbol))
	pprofGroup.GET("/symbol", gin.WrapF(pprof.Symbol))
	pprofGroup.GET("/trace", gin.WrapF(pprof.Trace))
	pprofGroup.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
	pprofGroup.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
	pprofGroup.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
	pprofGroup.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
	pprofGroup.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
	pprofGroup.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
}
