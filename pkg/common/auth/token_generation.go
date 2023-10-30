package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func (auth *Auth) GenerateToken(id int, username, email string, role Role) (string, error) {

	var (
		issuedAt  = time.Now()
		expiresAt = time.Now().Add(time.Hour * 24 * time.Duration(auth.sessionDurationDays))
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
	return token.SignedString([]byte(auth.secret))
}
