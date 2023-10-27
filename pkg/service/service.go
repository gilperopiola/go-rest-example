package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

var _ ServiceLayer = (*service)(nil)

type ServiceLayer interface {
	Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error)
	Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error)

	CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error)
	GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error)
	UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error)
	DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error)
	SearchUsers(searchUsersRequest requests.SearchUsersRequest) (responses.SearchUsersResponse, error)
	ChangePassword(changePasswordRequest requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error)

	CreateUserPost(createUserPostRequest requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error)
}

type service struct {
	config     *config.Config
	auth       auth.AuthI
	repository repository.RepositoryLayer
}

func New(repository repository.RepositoryLayer, auth auth.AuthI, config *config.Config) *service {
	return &service{
		repository: repository,
		auth:       auth,
		config:     config,
	}
}
