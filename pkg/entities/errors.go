package entities

import (
	"errors"
)

var (
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
)
