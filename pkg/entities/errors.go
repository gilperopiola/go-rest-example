package entities

import (
	"errors"
)

var (
	ErrAllFieldsRequired           = errors.New("all fields required")
	ErrPasswordsDontMatch          = errors.New("passwords don't match")
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
	ErrCreatingUser                = errors.New("error creating user")
)
