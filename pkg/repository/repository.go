package repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Repository struct {
	Database Database
}

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...QueryOption) (models.User, error)
	UserExists(email, username string, opts ...QueryOption) bool
	DeleteUser(id int) (models.User, error)
}

func NewRepository(database Database) *Repository {
	return &Repository{Database: database}
}

type QueryOption func(*string)

func WithoutDeleted(query *string) {
	*query += " AND deleted = false"
}
