package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"

	"github.com/jinzhu/gorm"
)

// CreateUser inserts a user. Table structure can be found on the models package
func (r *repository) CreateUser(user models.User) (models.User, error) {
	if err := r.database.db.Create(&user).Error; err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrCreatingUser)
	}
	return user, nil
}

// GetUser retrieves a user, if it exists
func (r *repository) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {
	query := "(id = ? OR username = ? OR email = ?)"
	for _, opt := range opts {
		opt(&query)
	}

	err := r.database.db.Preload("Details").Preload("Posts").Where(query, user.ID, user.Username, user.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, common.Wrap(err, customErrors.ErrUserNotFound)
		}
		return models.User{}, common.Wrap(err, customErrors.ErrUnknown)
	}

	return user, nil
}

// UpdateUser updates the fields that are not empty on the model
func (r *repository) UpdateUser(user models.User) (models.User, error) {
	if err := r.database.db.Model(&user).Update(&user).Error; err != nil {
		return models.User{}, common.Wrap(err, customErrors.ErrUpdatingUser)
	}
	return user, nil
}

// DeleteUser soft-deletes a user. if it is already deleted it throws an error
func (r *repository) DeleteUser(id int) (models.User, error) {

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

func (r *repository) SearchUsers(username string, page, perPage int, opts ...options.PreloadOption) (models.Users, error) {
	var users models.Users

	for _, opt := range opts {
		r.database.db = opt(r.database.db)
	}

	username = "%" + username + "%"
	if err := r.database.db.Where("username LIKE ?", username).Offset(page * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return models.Users{}, common.Wrap(err, customErrors.ErrSearchingUsers)
	}
	return users, nil
}

// UserExists checks if a user with username or email exists
func (r *repository) UserExists(username, email string, opts ...options.QueryOption) bool {
	query := "(username = ? OR email = ?)"
	for _, opt := range opts {
		opt(&query)
	}

	var count int64
	r.database.db.Model(&models.User{}).Where(query, username, email).Count(&count)
	return count > 0
}
