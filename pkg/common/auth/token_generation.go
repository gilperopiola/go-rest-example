package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Role string

const (
	AnyRole   Role = "any"
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     Role   `json:"role"`
	jwt.RegisteredClaims
}

/*----------------------------------------------------------------------------------------
// GenerateToken is called by the User Model on the Login endpoint.
------------------------*/

func GenerateToken(id int, username, email string, role Role, secret string, sessDurationDays int) (string, error) {

	var (
		issuedAt  = time.Now()
		expiresAt = time.Now().Add(time.Hour * 24 * time.Duration(sessDurationDays))
	)

	// Generate claims containing Username, Email, Role and ID
	claims := &CustomClaims{
		Username: username,
		Email:    email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprint(id),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	// Generate token (struct)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token (string)
	return token.SignedString([]byte(secret))
}
