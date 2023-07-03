package entities

import (
	"errors"
	"fmt"
)

var (
	// Generic
	ErrAllFieldsRequired = errors.New("all fields required")

	// Signup - Login
	ErrPasswordsDontMatch          = errors.New("passwords don't match")
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
	ErrInvalidEmailFormat          = errors.New("invalid email format")
	ErrCreatingUser                = errors.New("error creating user")
	ErrInvalidUsernameLength       = fmt.Errorf("username must be between %d and %d characters", USERNAME_MIN_LENGTH, USERNAME_MAX_LENGTH)
	ErrInvalidPasswordLength       = fmt.Errorf("password must be between %d and %d characters", PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH)
)
