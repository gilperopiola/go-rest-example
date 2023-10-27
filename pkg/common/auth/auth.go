package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthI interface {
	GenerateToken(user User) (string, error)
	ValidateToken(role Role, shouldMatchUserID bool) gin.HandlerFunc
}

func NewAuth(secret string, sessionDurationDays int) *Auth {
	return &Auth{
		secret:              secret,
		sessionDurationDays: sessionDurationDays,
	}
}

type Auth struct {
	secret              string
	sessionDurationDays int
}

type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
	jwt.RegisteredClaims
}
