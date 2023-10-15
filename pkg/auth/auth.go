package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func NewAuth(secret string, sessionDurationDays int) *Auth {
	return &Auth{
		secret:              secret,
		sessionDurationDays: sessionDurationDays,
	}
}

type AuthI interface {
	GenerateToken(user User, role Role) (string, error)
	ValidateToken(role Role, shouldMatchUserID bool) gin.HandlerFunc
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
