package auth

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	GenerateToken(user entities.User, role AuthRole) string
	ValidateToken() gin.HandlerFunc
	ValidateRole(role AuthRole) gin.HandlerFunc
	GetUserRole() AuthRole
	GetAdminRole() AuthRole
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
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Role     AuthRole `json:"role"`
	jwt.StandardClaims
}

type AuthRole string

const (
	UserRole  AuthRole = "user"
	AdminRole AuthRole = "admin"
)

func (auth *Auth) GetUserRole() AuthRole {
	return UserRole
}

func (auth *Auth) GetAdminRole() AuthRole {
	return AdminRole
}
