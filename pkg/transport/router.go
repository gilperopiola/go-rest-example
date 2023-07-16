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

func NewRouter(endpoints EndpointsIface, config config.Config) Router {
	var router Router
	router.Setup(endpoints, config)
	return router
}

/* ------------------- */

func (router *Router) Setup(endpoints EndpointsIface, config config.Config) {

	// Prepare router
	if !config.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Engine = gin.New()
	router.Use(getCORSConfig())

	// Set endpoints
	router.SetPublicEndpoints(endpoints)
	router.SetUserEndpoints(endpoints, config.JWT)
}

func (router *Router) SetPublicEndpoints(endpoints EndpointsIface) {
	public := router.Group("/")
	public.POST("/signup", endpoints.Signup)
	public.POST("/login", endpoints.Login)
}

func (router *Router) SetUserEndpoints(endpoints EndpointsIface, jwtConfig config.JWTConfig) {
	users := router.Group("/users", auth.ValidateToken(jwtConfig))
	users.GET("/:user_id", endpoints.GetUser)
	users.PATCH("/:user_id", endpoints.UpdateUser)
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
