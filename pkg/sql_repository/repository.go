package sql_repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

// Compile time check to ensure repository implements the RepositoryLayer interface
var _ RepositoryLayer = (*repository)(nil)

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...any) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	UpdatePassword(userID int, password string) error
	DeleteUser(user models.User) (models.User, error)
	SearchUsers(page, perPage int, opts ...any) (models.Users, error)
	CreateUserPost(post models.UserPost) (models.UserPost, error)
}

type repository struct {
	*Database
}

func New(database *Database) *repository {
	return &repository{Database: database}
}
