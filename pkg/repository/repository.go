package repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

type Repository struct {
	Database Database
}

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...options.QueryOption) (models.User, error)
	DeleteUser(id int) (models.User, error)
	UserExists(username, email string, opts ...options.QueryOption) bool

	CreateUserPost(post models.UserPost) (models.UserPost, error)
}

func NewRepository(database Database) *Repository {
	return &Repository{Database: database}
}
