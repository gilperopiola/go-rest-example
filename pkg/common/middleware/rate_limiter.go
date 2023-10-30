package middleware

import (
	"net/http"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func NewRateLimiterMiddleware(limiter *rate.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, common.HTTPResponse{
				Success: false,
				Content: nil,
				Error:   "too many requests",
			})
			return
		}
		c.Next()
	}
}

func NewRateLimiter(requestsPerSecond int) *rate.Limiter { // TODO RPS to Config var
	return rate.NewLimiter(rate.Every(time.Second/time.Duration(requestsPerSecond)), requestsPerSecond)
}
