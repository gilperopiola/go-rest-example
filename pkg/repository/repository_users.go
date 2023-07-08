package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/jinzhu/gorm"
)

func (r *Repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Create(&user).Error; err != nil {
		return models.User{}, utils.JoinErrors(ErrCreatingUser, err)
	}

	return user, nil
}

func (r *Repository) UpdateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Model(&user).Update(&user).Error; err != nil {
		return models.User{}, utils.JoinErrors(ErrUpdatingUser, err)
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

func (r *Repository) GetUser(user models.User) (models.User, error) {
	var databaseUser models.User

	if err := r.Database.DB.Where("id = ? OR username = ? OR email = ?", user.ID, user.Username, user.Email).First(&databaseUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, utils.JoinErrors(ErrGettingUser, err)
		}
		return models.User{}, utils.JoinErrors(ErrUnknown, err)
	}
	return databaseUser, nil
}
