package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type Repository struct {
	Database Database
}

type UserI interface {
	OverwriteFields(username, email, password string) models.User
	PasswordMatches(password string) bool
	ToEntity() entities.User
	GetAuthRole() entities.Role
}

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...utils.QueryOption) (models.User, error)
	UserExists(email, username string, opts ...utils.QueryOption) bool
	DeleteUser(id int) (models.User, error)
}

func NewRepository(database Database) *Repository {
	return &Repository{Database: database}
}

var (
	// - General errors
	ErrUnknown = errors.New("error unknown")

	// - User errors
	ErrCreatingUser       = errors.New("error creating user")
	ErrGettingUser        = errors.New("error getting user")
	ErrUpdatingUser       = errors.New("error updating user")
	ErrUserNotFound       = errors.New("error, user not found")
	ErrUserAlreadyDeleted = errors.New("error, user already deleted")
)
