package transport

import (
	"fmt"
	"io"
	"net/http/pprof"
	"os"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type router struct {
	*gin.Engine
}

func NewRouter(t TransportLayer, middlewares ...gin.HandlerFunc) router {

	// Create router. Set debug/release mode
	var router router
	router.prepare(!common.Cfg.Debug)

	// Add middlewares
	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	// Set endpoints
	router.setEndpoints(t)

	return router
}

func (router *router) prepare(isProd bool) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DefaultWriter = io.MultiWriter(common.Logger)      // Logger
	gin.DefaultErrorWriter = io.MultiWriter(common.Logger) // Logger

	router.Engine = gin.New()
	router.Engine.SetTrustedProxies(nil)
	if ok := binding.Validator.Engine().(*validator.Validate); ok == nil {
		fmt.Printf("error setting router validator")
		os.Exit(1)
	}
}

/*-----------------------------
//     Routes / Endpoints
//---------------------------*/

func (router *router) setEndpoints(transport TransportLayer) {

	// Standard endpoints
	router.GET("/health", transport.healthCheck)

	// V1
	v1 := router.Group("/v1")
	{
		router.setV1Endpoints(v1, transport)
	}

	// Monitoring
	if common.Cfg.Monitoring.PrometheusEnabled {
		router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	// Profiling
	if common.Cfg.Profiling {
		router.profiling()
	}

	fmt.Println("")
}

func (router *router) setV1Endpoints(v1 *gin.RouterGroup, transport TransportLayer) {

	// Auth
	v1.POST("/signup", transport.signup)
	v1.POST("/login", transport.login)

	// Users
	users := v1.Group("/users", auth.ValidateToken(auth.AnyRole, true, common.Cfg.Auth.JWTSecret))
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
	admin := v1.Group("/admin", auth.ValidateToken(auth.AdminRole, false, common.Cfg.Auth.JWTSecret))
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
