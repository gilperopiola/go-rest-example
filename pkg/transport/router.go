package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(transport TransportProvider, config config.ConfigInterface, auth auth.AuthInterface) Router {
	var router Router
	router.Setup(transport, config, auth)
	return router
}

/* ------------------- */

func (router *Router) Setup(transport TransportProvider, config config.ConfigInterface, auth auth.AuthInterface) {

	// Prepare router
	if !config.GetDebugMode() {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Engine = gin.New()
	router.Use(getCORSConfig())

	// Set endpoints
	router.SetPublicEndpoints(transport)
	router.SetUserEndpoints(transport, config.GetJWTConfig(), auth)
}

func (router *Router) SetPublicEndpoints(transport TransportProvider) {
	public := router.Group("/")
	public.POST("/signup", transport.Signup)
	public.POST("/login", transport.Login)
}

func (router *Router) SetUserEndpoints(transport TransportProvider, jwtConfig config.JWTConfig, auth auth.AuthInterface) {
	users := router.Group("/users", auth.ValidateToken())
	users.GET("/:user_id", transport.GetUser)
	users.PATCH("/:user_id", transport.UpdateUser)
	users.DELETE("/:user_id", transport.DeleteUser)
}

/* ------------------- */

func getCORSConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
