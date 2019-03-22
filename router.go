package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterActions interface {
	Setup()
}

type MyRouter struct {
	*gin.Engine
}

func (router *MyRouter) Setup() {
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
		public.POST("/SignUp", SignUp)
		public.POST("/LogIn", LogIn)
	}

	user := router.Group("/User", validateToken("User"))
	{
		user.GET("/Self", GetSelf)
	}

	admin := router.Group("/Admin", validateToken("Admin"))
	{
		user := admin.Group("/User")
		{
			user.POST("", CreateUser)
			user.GET("/:id", ReadUser)
			user.PUT("/:id", UpdateUser)
		}
		admin.GET("/Users", ReadUsers)

		movie := admin.Group("/Movie")
		{
			movie.POST("", CreateMovie)
			movie.GET("/:id", ReadMovie)
			movie.PUT("/:id", UpdateMovie)
		}
		admin.GET("/Movies", ReadMovies)

		director := admin.Group("/Director")
		{
			director.POST("", CreateDirector)
			director.GET("/:id", ReadDirector)
			director.PUT("/:id", UpdateDirector)
		}
		admin.GET("/Directors", ReadDirectors)

		actor := admin.Group("/Actor")
		{
			actor.POST("", CreateActor)
			actor.GET("/:id", ReadActor)
			actor.PUT("/:id", UpdateActor)
		}
		admin.GET("/Actors", ReadActors)
	}
}
