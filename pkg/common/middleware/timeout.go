package middleware

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func NewTimeoutMiddleware(timeoutSeconds int) gin.HandlerFunc {
	return func() gin.HandlerFunc {
		return timeout.New(
			timeout.WithTimeout(time.Duration(timeoutSeconds) * time.Second),
		)
	}()
}
