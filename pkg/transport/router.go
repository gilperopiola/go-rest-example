package transport

import (
	"fmt"
	"io"
	"net/http/pprof"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type router struct {
	*gin.Engine
}

func NewRouter(t TransportLayer, middlewares ...gin.HandlerFunc) router {
	var router router
	router.prepare(!common.Cfg.Debug)

	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	router.setEndpoints(t)

	return router
}

func (router *router) prepare(isProd bool) {

	// Set Prod / Debug mode
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set logger
	gin.DefaultWriter = io.MultiWriter(common.Logger)
	gin.DefaultErrorWriter = io.MultiWriter(common.Logger)

	// Create gin Engine
	router.Engine = gin.New()
	router.Engine.SetTrustedProxies(nil)
}

/*-----------------------------
//     Routes / Endpoints
//---------------------------*/

func (router *router) setEndpoints(transport TransportLayer) {

	router.GET("/health", transport.healthCheck)

	v1 := router.Group("/v1")
	router.setV1Endpoints(v1, transport)

	if common.Cfg.Monitoring.PrometheusEnabled {
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	if common.Cfg.Profiling {
		router.profiling()
	}

	fmt.Println("")
}

func (router *router) setV1Endpoints(v1 *gin.RouterGroup, transport TransportLayer) {

	v1.POST("/signup", transport.signup)
	v1.POST("/login", transport.login)

	// Users
	users := v1.Group("/users", auth.ValidateToken(auth.AnyRole, true, common.Cfg.JWTSecret))
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
	admin := v1.Group("/admin", auth.ValidateToken(auth.AdminRole, false, common.Cfg.JWTSecret))
	{
		admin.POST("/user", transport.createUser)
		admin.GET("/users", transport.searchUsers)
	}
}

// Profiling, only called if the config is set to true
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
