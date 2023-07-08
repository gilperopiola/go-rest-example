package entities

import (
	"errors"
	"fmt"
)

var (

	// Generic

	ErrUnauthorized      = errors.New("unauthorized")
	ErrBindingRequest    = errors.New("error binding request")
	ErrAllFieldsRequired = errors.New("all fields required")

	// Signup - Login - Users in general

	ErrCreatingUser = errors.New("error creating user")
	ErrUserNotFound = errors.New("user not found")

	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrInvalidUsernameLength = fmt.Errorf("username must be between %d and %d characters", USERNAME_MIN_LENGTH, USERNAME_MAX_LENGTH)
	ErrInvalidPasswordLength = fmt.Errorf("password must be between %d and %d characters", PASSWORD_MIN_LENGTH, PASSWORD_MAX_LENGTH)

	ErrPasswordsDontMatch          = errors.New("passwords don't match")
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
	ErrWrongPassword               = errors.New("wrong password")
)
