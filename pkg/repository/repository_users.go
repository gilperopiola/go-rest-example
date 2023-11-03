package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"

	"gorm.io/gorm"
)

// CreateUser inserts a user. Table structure can be found on the models package
func (r *repository) CreateUser(user models.User) (models.User, error) {
	db := r.database.DB()
	if err := db.Create(&user).Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrCreatingUser)
	}
	return user, nil
}

// GetUser retrieves a user, if it exists
func (r *repository) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {
	db := r.database.DB()

	// Query by ID, username or email
	query := "(id = ? OR username = ? OR email = ?)"

	// WithoutDeleted, WithDetails, WithPosts
	for _, opt := range opts {
		db = opt(db, &query)
	}

	// Get user
	if err := db.Where(query, user.ID, user.Username, user.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, common.Wrap(err.Error(), common.ErrUserNotFound)
		}
		return models.User{}, common.Wrap(err.Error(), common.ErrUnknown)
	}

	return user, nil
}

// UpdateUser updates the fields that are not empty on the model
func (r *repository) UpdateUser(user models.User) (models.User, error) {
	db := r.database.DB()
	tx := db.Begin()

	// Update user
	if err := tx.Omit("Details").Save(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, common.Wrap(err.Error(), common.ErrUpdatingUser)
	}

	// Update user details
	if user.Details.ID != 0 {
		if err := tx.Save(&user.Details).Error; err != nil {
			tx.Rollback()
			return models.User{}, common.Wrap(err.Error(), common.ErrUpdatingUserDetail)
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrInDBTransaction)
	}

	return user, nil
}

// DeleteUser soft-deletes a user. if it is already deleted it throws an error
func (r *repository) DeleteUser(id int) (models.User, error) {
	var err error
	user := models.User{ID: id}

	// First, retrieve the user
	if user, err = r.GetUser(user); err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrGettingUser)
	}

	// If it's already deleted, return an error
	if user.Deleted {
		return models.User{}, common.ErrUserAlreadyDeleted
	}

	// Then, mark the user as deleted and save it
	user.Deleted = true
	if _, err := r.UpdateUser(user); err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrUpdatingUser)
	}

	return user, nil
}

func (r *repository) SearchUsers(page, perPage int, opts ...options.QueryOption) (models.Users, error) {
	db := r.database.DB()
	var users models.Users

	// WithUsername, WithDetails, WithPosts, WithoutDeleted
	for _, opt := range opts {
		db = opt(db, nil)
	}

	if err := db.Offset(page * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return models.Users{}, common.Wrap(err.Error(), common.ErrSearchingUsers)
	}

	return users, nil
}

// UserExists checks if a user with username or email exists
func (r *repository) UserExists(username, email string, opts ...options.QueryOption) bool {
	db := r.database.DB()
	query := "(username = ? OR email = ?)"

	// WithoutDeleted
	for _, opt := range opts {
		db = opt(db, &query)
	}

	var count int64
	db.Model(&models.User{}).Where(query, username, email).Count(&count)
	return count > 0
}

// CreateUserPost inserts a new post on the database. Title is required, body is optional
func (r *repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	db := r.database.DB()
	if err := db.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err.Error(), common.ErrCreatingUserPost)
	}
	return post, nil
}
