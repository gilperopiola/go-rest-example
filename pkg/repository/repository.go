package repository

import (
	"errors"
	"fmt"

	models "github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/jinzhu/gorm"
)

type Repositorier interface {
	CreateUser(user models.User) (models.User, error)
	UserExists(email, username string) bool
}

type Repository struct {
	Database Database
}

func (r *Repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Create(&user).Error; err != nil {
		return models.User{}, fmt.Errorf("%w:%w", ErrCreatingUser, err)
	}

	return user, nil
}

func (r *Repository) UserExists(email, username string) bool {
	var user models.User

	if err := r.Database.DB.Where("email = ? OR username = ?", email, username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}
