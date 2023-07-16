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

	if err := r.Database.DB.Where("(id = ? OR username = ? OR email = ?) AND deleted = false", user.ID, user.Username, user.Email).First(&databaseUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, utils.JoinErrors(ErrGettingUser, err)
		}
		return models.User{}, utils.JoinErrors(ErrUnknown, err)
	}
	return databaseUser, nil
}

func (r *Repository) DeleteUser(id int) (models.User, error) {
	var user models.User

	// First, retrieve the user
	if err := r.Database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, utils.JoinErrors(ErrGettingUser, err)
		}
		return models.User{}, utils.JoinErrors(ErrUnknown, err)
	}

	// If it's already deleted, return an error
	if user.Deleted {
		return models.User{}, ErrUserAlreadyDeleted
	}

	// Then, mark the user as deleted and save it
	user.Deleted = true
	if err := r.Database.DB.Save(&user).Error; err != nil {
		return models.User{}, utils.JoinErrors(ErrUpdatingUser, err)
	}

	return user, nil
}
