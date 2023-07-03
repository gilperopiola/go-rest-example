package repository

import models "github.com/gilperopiola/go-rest-example/pkg/models"

type Repository interface {
	CreateUser(user *models.User) error
	UserExists(email, username string) bool
}

type RepositoryHandler struct {
	Database *Database
}

func (r *RepositoryHandler) CreateUser(user *models.User) error {
	if err := r.Database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *RepositoryHandler) UserExists(email, username string) bool {
	return false
}
