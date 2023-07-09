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

func (router *Router) Setup(endpoints EndpointsIface, jwtConfig config.JWTConfig) {

	// Prepare router
	gin.SetMode(gin.DebugMode)
	router.Engine = gin.New()
	router.Use(getCORSConfig())

	// Public endpoints
	public := router.Group("/")
	public.POST("/signup", endpoints.Signup)
	public.POST("/login", endpoints.Login)

	// Private endpoints
	users := router.Group("/users", auth.ValidateToken(jwtConfig))
	users.GET("/:user_id", endpoints.GetUser)
	users.PATCH("/:user_id", endpoints.UpdateUser)
}

func getCORSConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
