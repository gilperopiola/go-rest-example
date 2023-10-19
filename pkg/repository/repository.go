package repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

// Compile time check to ensure repository implements the RepositoryLayer interface
var _ RepositoryLayer = (*repository)(nil)

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...options.QueryOption) (models.User, error)
	DeleteUser(id int) (models.User, error)
	UserExists(username, email string, opts ...options.QueryOption) bool

	CreateUserPost(post models.UserPost) (models.UserPost, error)
}

type repository struct {
	Database Database
}

func New(database Database) *repository {
	return &repository{Database: database}
}
