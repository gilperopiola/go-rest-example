package auth

import (
	"fmt"
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/entities"
	"github.com/golang-jwt/jwt/v4"
)

func (auth *Auth) GenerateToken(user entities.User, role entities.Role) (string, error) {

	var (
		// Session duration can be set in the config
		issuedAt  = time.Now()
		expiresAt = time.Now().Add(time.Hour * 24 * time.Duration(auth.sessionDurationDays))
	)

	// Generate claims containing Username, Email, Role and ID
	claims := &CustomClaims{
		Username: user.Username,
		Email:    user.Email,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        fmt.Sprint(user.ID),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	// Generate token (struct)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token (string)
	return token.SignedString([]byte(auth.secret))
}
