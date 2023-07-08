package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Repository struct {
	Database Database
}

type RepositoryIFace interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User) (models.User, error)
	UserExists(email, username string) bool
}

var (

	// General errors
	ErrUnknown = errors.New("error unknown")

	// User errors
	ErrCreatingUser = errors.New("error creating user")
	ErrUpdatingUser = errors.New("error updating user")
	ErrGettingUser  = errors.New("error getting user")
)
