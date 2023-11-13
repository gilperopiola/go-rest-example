package middleware

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func NewRateLimiterMiddleware(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.Allow() {
			c.Next()
		}
		c.Error(common.ErrTooManyRequests)
		c.Abort()
	}
}

func NewRateLimiter(requestsPerSecond int) *rate.Limiter {
	return rate.NewLimiter(rate.Every(time.Second/time.Duration(requestsPerSecond)), requestsPerSecond)
}
