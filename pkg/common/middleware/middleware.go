package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type Middlewares []gin.HandlerFunc

func NewTimeoutMiddleware(timeoutSeconds int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(timeoutSeconds)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
	)
}

func NewLoggerToContextMiddleware(logger LoggerI) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("logger", logger)
		c.Next()
	}
}

func NewCORSConfigMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authentication", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Authentication", "Authorization", "Content-Type"},
	})
}
