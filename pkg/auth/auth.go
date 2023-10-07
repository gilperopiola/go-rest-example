package auth

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	GenerateToken(user entities.User, role entities.Role) string
	ValidateToken(role entities.Role) gin.HandlerFunc
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
	jwt.StandardClaims
}
