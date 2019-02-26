package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

type Config struct {
	PORT     string
	DEBUG    bool
	DATABASE struct {
		TYPE     string
		USERNAME string
		PASSWORD string
		HOSTNAME string
		PORT     string
		SCHEMA   string
	}
	JWT struct {
		SECRET           string
		SESSION_DURATION int
	}
	USERS struct {
		USERNAME_MIN_CHARACTERS int
		USERNAME_MAX_CHARACTERS int
	}
}

var config *Config
var db *gorm.DB
var router *gin.Engine

func main() {
	setupConfig()
	setupDatabase()
	setupRouter()

	defer db.Close()
	log.Println("server started")
	router.Run(":" + config.PORT)
}

func setupRouter() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.New()

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

	user := router.Group("/User", validateToken("Admin"))
	{
		user.POST("", CreateUser)
		user.GET("/:id", ReadUser)
		user.PUT("/:id", UpdateUser)
	}
	router.GET("/Users", validateToken("Admin"), ReadUsers)

	movie := router.Group("/Movie", validateToken("Admin"))
	{
		movie.POST("", CreateMovie)
		movie.GET("/:id", ReadMovie)
		movie.PUT("/:id", UpdateMovie)
	}
	router.GET("/Movies", validateToken("Admin"), ReadMovies)

	director := router.Group("/Director", validateToken("Admin"))
	{
		director.POST("", CreateDirector)
		director.GET("/:id", ReadDirector)
		director.PUT("/:id", UpdateDirector)
	}
	router.GET("/Directors", validateToken("Admin"), ReadDirectors)

}

func setupConfig() {
	viper.SetConfigName("config") //config filename without the .JSON or .YAML extension
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}
}
