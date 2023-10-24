package repository

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"

	"github.com/jinzhu/gorm"
)

// CreateUser inserts a user. Table structure can be found on the models package
func (r *repository) CreateUser(user models.User) (models.User, error) {
	if err := r.database.db.Create(&user).Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrCreatingUser)
	}
	return user, nil
}

// GetUser retrieves a user, if it exists
func (r *repository) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {

	// get by id, username or email
	query := "(id = ? OR username = ? OR email = ?)"

	// only non deleted users
	for _, opt := range opts {
		opt(&query)
	}

	// preload user details and posts
	err := r.database.db.Preload("Details").Preload("Posts").Where(query, user.ID, user.Username, user.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, common.Wrap(err.Error(), common.ErrUserNotFound)
		}
		return models.User{}, common.Wrap(err.Error(), common.ErrUnknown)
	}

	return user, nil
}

// UpdateUser updates the fields that are not empty on the model
func (r *repository) UpdateUser(user models.User) (models.User, error) {
	if err := r.database.db.Model(&user).Update(&user).Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrUpdatingUser)
	}
	return user, nil
}

// DeleteUser soft-deletes a user. if it is already deleted it throws an error
func (r *repository) DeleteUser(id int) (models.User, error) {

	// first, retrieve the user
	user := models.User{ID: id}
	var err error
	if user, err = r.GetUser(user); err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrGettingUser)
	}

	// if it's already deleted, return an error
	if user.Deleted {
		return models.User{}, common.ErrUserAlreadyDeleted
	}

	// then, mark the user as deleted and save it
	user.Deleted = true
	if _, err := r.UpdateUser(user); err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrUpdatingUser)
	}

	return user, nil
}

func (r *repository) SearchUsers(username string, page, perPage int, opts ...options.PreloadOption) (models.Users, error) {
	var users models.Users

	// preload user details and posts
	for _, opt := range opts {
		r.database.db = opt(r.database.db)
	}

	// if username is provided, apply the filter
	if username != "" {
		searchPattern := "%" + username + "%"
		r.database.db = r.database.db.Where("username LIKE ?", searchPattern)
	}

	if err := r.database.db.Offset(page * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return models.Users{}, common.Wrap(err.Error(), common.ErrSearchingUsers)
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

// CreateUserPost inserts a new post on the database. Title is required, body is optional
func (r *repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	if err := r.database.db.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err.Error(), common.ErrCreatingUserPost)
	}
	return post, nil
}
