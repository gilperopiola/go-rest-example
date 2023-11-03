package repository

import (
	"log"

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
	SearchUsers(page, perPage int, opts ...options.QueryOption) (models.Users, error)
	UserExists(username, email string, opts ...options.QueryOption) bool

	CreateUserPost(post models.UserPost) (models.UserPost, error)
}

func New(database database) *repository {
	log.Println("Repository OK")
	return &repository{database: database}
}

type repository struct {
	database database
}
