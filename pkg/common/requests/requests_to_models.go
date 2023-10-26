package requests

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

//-------------------------------
//     REQUEST TO USER MODEL
//-------------------------------

func (r *SignupRequest) ToUserModel() models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *LoginRequest) ToUserModel() models.User {
	user := models.User{Password: r.Password}

	if validEmailRegex.MatchString(r.UsernameOrEmail) {
		user.Email = r.UsernameOrEmail
	} else {
		user.Username = r.UsernameOrEmail
	}

	return user
}

func (r *CreateUserRequest) ToUserModel() models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:   r.IsAdmin,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (r *GetUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (r *UpdateUserRequest) ToUserModel() models.User {
	firstName, lastName := "", ""

	if r.FirstName != nil {
		firstName = *r.FirstName
	}
	if r.LastName != nil {
		lastName = *r.LastName
	}

	return models.User{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: firstName,
			LastName:  lastName,
		},
	}
}

func (r *DeleteUserRequest) ToUserModel() models.User {
	return models.User{ID: r.ID}
}

func (r *SearchUsersRequest) ToUserModel() models.User {
	return models.User{Username: r.Username}
}

//--------------------------------
//   REQUEST TO USER POST MODEL
//--------------------------------

func (r *CreateUserPostRequest) ToUserPostModel() models.UserPost {
	return models.UserPost{
		UserID: r.UserID,
		Title:  r.Title,
		Body:   r.Body,
	}
}
