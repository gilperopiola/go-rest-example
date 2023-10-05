package transport

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func addLoggerToContextMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func addCorsConfigMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
