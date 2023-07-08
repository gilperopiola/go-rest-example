package repository

import "errors"

var (

	// General errors
	ErrUnknown = errors.New("error unknown. check the logs for the error trace")

	// User errors
	ErrCreatingUser = errors.New("error creating user")
	ErrGettingUser  = errors.New("error getting user")
)
