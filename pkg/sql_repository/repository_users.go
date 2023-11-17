package sql_repository

import (
	"errors"
	"strings"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/sql_repository/options"

	"gorm.io/gorm"
)

func handleUserError(err error, defaultErr *common.Error) error {
	errMsg := err.Error()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return common.Wrap(errMsg, common.ErrUserNotFound)
	}
	if strings.Contains(errMsg, "Error 1062") { // Duplicate entry for key
		return common.Wrap(errMsg, common.ErrUsernameOrEmailAlreadyInUse)
	}

	return common.Wrap(errMsg, defaultErr)
}

/*-------------------------
//      Create User
//-----------------------*/

func (r *repository) CreateUser(user models.User) (models.User, error) {
	if err := r.Database.Create(&user).Error; err != nil {
		return models.User{}, handleUserError(err, common.ErrCreatingUser)
	}
	return user, nil
}

/*-------------------------
//       Get User
//-----------------------*/

func (r *repository) GetUser(user models.User, opts ...any) (models.User, error) {

	// Query by ID, username or email.
	// If we have the ID, discard the other fields.
	// If we have Username, discard Email and viceversa.
	query := "(id = ? OR username = ? OR email = ?)"
	if user.ID != 0 {
		user.Username, user.Email = "", ""
	}
	if user.Username != "" {
		user.Email = ""
	}
	if user.Email != "" {
		user.Username = ""
	}

	// WithoutDeleted, WithDetails, WithPosts
	db := options.ApplyQueryOptions(r.Database.DB, &query, opts...)

	// Get user
	if err := db.Where(query, user.ID, user.Username, user.Email).First(&user).Error; err != nil {
		return models.User{}, handleUserError(err, common.ErrGettingUser)
	}

	return user, nil
}

/*--------------------------------------------
//      Update User -> (non-empty fields)
//-----------------------*/

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	tx := r.Database.DB.Begin()

	// Update user
	if err := tx.Omit("Details").Save(&user).Error; err != nil {
		tx.Rollback()
		return models.User{}, handleUserError(err, common.ErrUpdatingUser)
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

/*-------------------------
//    Update Password
//-----------------------*/

func (r *repository) UpdatePassword(userID int, newPassword string) error {
	if err := r.Database.Model(&models.User{}).Where("id = ?", userID).Update("password", newPassword).Error; err != nil {
		return common.Wrap(err.Error(), common.ErrUpdatingUser)
	}
	return nil
}

/*--------------------------------------
//      Delete User -> (soft-delete)
//-----------------------*/

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	if err := r.Database.Model(&user).Update("deleted", true).Error; err != nil {
		return models.User{}, common.Wrap(err.Error(), common.ErrDeletingUser)
	}
	return user, nil
}

/*-------------------------
//     Search Users
//-----------------------*/

func (r *repository) SearchUsers(page, perPage int, opts ...any) (models.Users, error) {

	// WithUsername, WithDetails, WithPosts, WithoutDeleted
	db := options.ApplyQueryOptions(r.Database.DB, nil, opts...)

	var users models.Users
	if err := db.Offset(page * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return models.Users{}, common.Wrap(err.Error(), common.ErrSearchingUsers)
	}

	return users, nil
}

/*-------------------------
//    Create User Post
//-----------------------*/

func (r *repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	if err := r.Database.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err.Error(), common.ErrCreatingUserPost)
	}
	return post, nil
}
