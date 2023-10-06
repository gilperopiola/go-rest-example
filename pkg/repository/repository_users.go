package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/jinzhu/gorm"
)

// CreateUser creates a user on the database. Id, username and email are unique
func (r *Repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Create(&user).Error; err != nil {
		return models.User{}, utils.WrapErrors(err, ErrCreatingUser)
	}

	return user, nil
}

// UpdateUser updates the user on the database, skipping fields that are empty
func (r *Repository) UpdateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Model(&user).Update(&user).Error; err != nil {
		return models.User{}, utils.WrapErrors(err, ErrUpdatingUser)
	}

	return user, nil
}

// UserExists checks if a user exists on the database
func (r *Repository) UserExists(email, username string, onlyNonDeleted bool) bool {
	var user models.User

	query := buildNonDeletedQuery("(email = ? OR username = ?)", onlyNonDeleted)

	if err := r.Database.DB.Where(query, email, username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

// GetUser retrieves a user from the database, if it exists
func (r *Repository) GetUser(user models.User, onlyNonDeleted bool) (models.User, error) {
	query := buildNonDeletedQuery("(id = ? OR username = ? OR email = ?)", onlyNonDeleted)

	err := r.Database.DB.Where(query, user.ID, user.Username, user.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, utils.WrapErrors(err, ErrGettingUser)
		}
		return models.User{}, utils.WrapErrors(err, ErrUnknown)
	}

	return user, nil
}

// DeleteUser marks a user as deleted on the database, if it is already deleted it throws an error
func (r *Repository) DeleteUser(id int) (user models.User, err error) {

	// First, retrieve the user
	user.ID = id
	if user, err = r.GetUser(user, false); err != nil {
		return models.User{}, err
	}

	// If it's already deleted, return an error
	if user.Deleted {
		return models.User{}, ErrUserAlreadyDeleted
	}

	// Then, mark the user as deleted and save it
	user.Deleted = true
	if _, err := r.UpdateUser(user); err != nil {
		return models.User{}, utils.WrapErrors(err, ErrUpdatingUser)
	}

	return user, nil
}

func buildNonDeletedQuery(query string, onlyNonDeleted bool) string {
	if onlyNonDeleted {
		query += " AND deleted = false"
	}
	return query
}
