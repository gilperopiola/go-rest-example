package repository

import (
	"errors"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"

	"gorm.io/gorm"
)

/*-------------------------
//      Create User
//-----------------------*/

func (r *repository) CreateUser(user models.User) (models.User, error) {
	if err := r.database.DB().Create(&user).Error; err != nil {
		return models.User{}, handleCreateUserError(err)
	}
	return user, nil
}

func handleCreateUserError(err error) error {
	if strings.Contains(err.Error(), "Error 1062") { // Duplicate entry for key
		return common.Wrap(err.Error(), common.ErrUsernameOrEmailAlreadyInUse)
	}
	return common.Wrap(err.Error(), common.ErrCreatingUser)
}

/*-------------------------
//       Get User
//-----------------------*/

func (r *repository) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {
	db := r.database.DB()

	// Query by ID, username or email
	query := "(id = ? OR username = ? OR email = ?)"

	// If we have the ID, discard the other fields
	if user.ID != 0 {
		user.Username, user.Email = "", ""
	}

	// WithoutDeleted, WithDetails, WithPosts
	for _, opt := range opts {
		db = opt(db, &query)
	}

	// Get user
	if err := db.Where(query, user.ID, user.Username, user.Email).First(&user).Error; err != nil {
		return models.User{}, handleGetUserError(err)
	}

	return user, nil
}

func handleGetUserError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Wrap(err.Error(), common.ErrUserNotFound)
	}
	return common.Wrap(err.Error(), common.ErrGettingUser)
}

/*--------------------------------------------
//      Update User -> (non-empty fields)
//-----------------------*/

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	db := r.database.DB()
	tx := db.Begin()

	// Update user
	if err := tx.Omit("Details").Save(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, handleUpdateUserError(err)
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

func handleUpdateUserError(err error) error {
	if strings.Contains(err.Error(), "Error 1062") { // Duplicate entry for key
		return common.Wrap(err.Error(), common.ErrUsernameOrEmailAlreadyInUse)
	}
	return common.Wrap(err.Error(), common.ErrUpdatingUser)
}

/*-------------------------
//    Update Password
//-----------------------*/

func (r *repository) UpdatePassword(userID int, newPassword string) error {
	db := r.database.DB()

	if err := db.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error; err != nil {
		return common.Wrap(err.Error(), common.ErrUpdatingUser)
	}

	return nil
}

/*--------------------------------------
//      Delete User -> (soft-delete)
//-----------------------*/

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	var db = r.database.DB()

	if err := db.Model(&user).Update("deleted", true).Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrDeletingUser)
	}

	return user, nil
}

/*-------------------------
//     Search Users
//-----------------------*/

func (r *repository) SearchUsers(page, perPage int, opts ...options.QueryOption) (models.Users, error) {
	var db = r.database.DB()
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

/*-------------------------
//    Create User Post
//-----------------------*/

// Title is required, body is optional
func (r *repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	db := r.database.DB()
	if err := db.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err.Error(), common.ErrCreatingUserPost)
	}
	return post, nil
}
