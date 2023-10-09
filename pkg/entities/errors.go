package entities

import (
	"errors"
	"fmt"
)

// These are the service errors that can be returned by the service layer.
//
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go
// - IMPORTANT: If you add a new error, make sure to add it to the errorsMapToHTTPCode map in pkg/transport/errors_mapper.go

var (

	// - Generic

	ErrUnauthorized      = errors.New("unauthorized")
	ErrBindingRequest    = errors.New("binding request failed")
	ErrAllFieldsRequired = errors.New("all fields required")
	ErrNilError          = errors.New("nil error, unexpected behavior")
	ErrUnknown           = errors.New("unknown")

	// - Auth & Users

	ErrCreatingUser = errors.New("failed to create user")
	ErrUserNotFound = errors.New("user not found")

	ErrInvalidEmailFormat    = errors.New("invalid email format")
	ErrInvalidUsernameLength = fmt.Errorf("username must be between %d and %d characters", usernameMinLength, usernameMaxLength)
	ErrInvalidPasswordLength = fmt.Errorf("password must be between %d and %d characters", passwordMinLength, passwordMaxLength)

	ErrPasswordsDontMatch          = errors.New("passwords don't match")
	ErrUsernameOrEmailAlreadyInUse = errors.New("username or email already in use")
	ErrWrongPassword               = errors.New("wrong password")
)
