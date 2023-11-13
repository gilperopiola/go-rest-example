package requests

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

/*---------------------------------------------------------------------------
// modelDeps is used to inject the Config and Repository into the Models.
------------------------*/

func modelDeps(config *config.Config, repository models.RepositoryI) *models.ModelDependencies {
	return &models.ModelDependencies{
		Config:     config,
		Repository: repository,
	}
}

/*---------------
//    Signup
//-------------*/

func (r *SignupRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		Username: r.Username,
		Email:    r.Email,
		Password: common.Hash(r.Password, config.Auth.HashSalt),
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:           false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		ModelDependencies: modelDeps(config, repository),
	}
}

/*--------------
//    Login
//------------*/

func (r *LoginRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	if validEmailRegex.MatchString(r.UsernameOrEmail) {
		return models.User{
			Email:             r.UsernameOrEmail,
			Password:          r.Password,
			ModelDependencies: modelDeps(config, repository),
		}
	} else {
		return models.User{
			Username:          r.UsernameOrEmail,
			Password:          r.Password,
			ModelDependencies: modelDeps(config, repository),
		}
	}
}

/*---------------------
//    Create User
--------------------*/

func (r *CreateUserRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		Email:    r.Email,
		Username: r.Username,
		Password: common.Hash(r.Password, config.Auth.HashSalt),
		Deleted:  false,
		Details: models.UserDetail{
			FirstName: r.FirstName,
			LastName:  r.LastName,
		},
		IsAdmin:           r.IsAdmin,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		ModelDependencies: modelDeps(config, repository),
	}
}

/*--------------------
//     Get User
//------------------*/

func (r *GetUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		ModelDependencies: modelDeps(nil, repository),
	}
}

/*--------------------
//    Update User
//------------------*/

func (r *UpdateUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	firstName, lastName := "", ""
	if r.FirstName != nil {
		firstName = *r.FirstName
	}
	if r.LastName != nil {
		lastName = *r.LastName
	}

	return models.User{
		ID:       r.UserID,
		Username: r.Username,
		Email:    r.Email,
		Details: models.UserDetail{
			FirstName: firstName,
			LastName:  lastName,
		},
		ModelDependencies: modelDeps(nil, repository),
	}
}

/*--------------------
//    Delete User
//------------------*/

func (r *DeleteUserRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		ModelDependencies: modelDeps(nil, repository),
	}
}

/*--------------------
//    Search Users
//------------------*/

func (r *SearchUsersRequest) ToUserModel(repository models.RepositoryI) models.User {
	return models.User{
		Username:          r.Username,
		ModelDependencies: modelDeps(nil, repository),
	}
}

/*-----------------------
//    Change Password
//---------------------*/

func (r *ChangePasswordRequest) ToUserModel(config *config.Config, repository models.RepositoryI) models.User {
	return models.User{
		ID:                r.UserID,
		Password:          r.OldPassword,
		ModelDependencies: modelDeps(config, repository),
	}
}

/*------------------------
//    Create User Post
//----------------------*/

func (r *CreateUserPostRequest) ToUserPostModel(repository models.RepositoryI) models.UserPost {
	return models.UserPost{
		UserID:            r.UserID,
		Title:             r.Title,
		Body:              r.Body,
		ModelDependencies: modelDeps(nil, repository),
	}
}
