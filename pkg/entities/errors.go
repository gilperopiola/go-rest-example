package entities

import (
	"errors"
)

var (
	// Signup
	ErrAllFieldsRequired           = errors.New("all fields required")
	ErrPasswordsDontMatch          = errors.New("passwords don't match")
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
	ErrInvalidEmailFormat          = errors.New("invalid email format")
	ErrCreatingUser                = errors.New("error creating user")
)
