package transport

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterIFace interface {
	Setup()
}

type Router struct {
	*gin.Engine
}

func (router *Router) Setup(endpoints EndpointsIface) {
	gin.SetMode(gin.DebugMode)
	router.Engine = gin.New()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Authentication", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	}))

	public := router.Group("/")
	{
		public.POST("/signup", endpoints.Signup)
		public.POST("/login", endpoints.Login)
	}

	//user := router.Group("/User", validateToken("User"))
	//{
	//	user.GET("/Self", GetSelf)
	//}
	//
	//admin := router.Group("/Admin", validateToken("Admin"))
	//{
	//	user := admin.Group("/User")
	//	{
	//		user.POST("", CreateUser)
	//		user.GET("/:id", ReadUser)
	//		user.PUT("/:id", UpdateUser)
	//	}
	//	admin.GET("/Users", ReadUsers)
	//}
}
