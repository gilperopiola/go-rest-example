package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"github.com/jinzhu/gorm"
)

// CreateUser creates a user on the database. Id, username and email are unique
func (r *Repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Create(&user).Error; err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrCreatingUser)
	}
	return user, nil
}

// GetUser retrieves a user from the database, if it exists
func (r *Repository) GetUser(user models.User, opts ...QueryOption) (models.User, error) {
	query := "(id = ? OR username = ? OR email = ?)"
	for _, opt := range opts {
		opt(&query)
	}

	err := r.Database.DB.Preload("Details").Preload("Posts").Where(query, user.ID, user.Username, user.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, common.Wrap(err, customErrors.ErrUserNotFound)
		}
		return models.User{}, common.Wrap(err, customErrors.ErrUnknown)
	}

	return user, nil
}

// UpdateUser updates the user on the database, skipping fields that are empty
func (r *Repository) UpdateUser(user models.User) (models.User, error) {
	if err := r.Database.DB.Model(&user).Update(&user).Error; err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrUpdatingUser)
	}
	return user, nil
}

// DeleteUser marks a user as deleted on the database, if it is already deleted it throws an error
func (r *Repository) DeleteUser(id int) (models.User, error) {

	// First, retrieve the user
	user := models.User{ID: id}
	var err error
	if user, err = r.GetUser(user); err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrGettingUser)
	}

	// If it's already deleted, return an error
	if user.Deleted {
		return models.User{}, customErrors.ErrUserAlreadyDeleted
	}

	// Then, mark the user as deleted and save it
	user.Deleted = true
	if _, err := r.UpdateUser(user); err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrUpdatingUser)
	}

	return user, nil
}

// UserExists checks if a user exists on the database
func (r *Repository) UserExists(email, username string, opts ...QueryOption) bool {
	query := "(email = ? OR username = ?)"
	for _, opt := range opts {
		opt(&query)
	}

	var count int64
	r.Database.DB.Model(&models.User{}).Where(query, email, username).Count(&count)
	return count > 0
}

func (r *Repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	if err := r.Database.DB.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err, customErrors.ErrCreatingUserPost)
	}
	return post, nil
}
