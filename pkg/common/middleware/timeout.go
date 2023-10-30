package middleware

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func NewTimeoutMiddleware(timeoutSeconds int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(timeoutSeconds)*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
	)
}
