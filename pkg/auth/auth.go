package auth

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthI interface {
	GenerateToken(user entities.User, role entities.Role) (string, error)
	ValidateToken(role entities.Role, shouldMatchUserID bool) gin.HandlerFunc
}

type Auth struct {
	secret              string
	sessionDurationDays int
}

func NewAuth(secret string, sessionDurationDays int) *Auth {
	return &Auth{
		secret:              secret,
		sessionDurationDays: sessionDurationDays,
	}
}

type CustomClaims struct {
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Role     entities.Role `json:"role"`
	jwt.RegisteredClaims
}
