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
	Signup(request *requests.SignupRequest) (responses.SignupResponse, error)
	Login(request *requests.LoginRequest) (responses.LoginResponse, error)

	CreateUser(request *requests.CreateUserRequest) (responses.CreateUserResponse, error)
	GetUser(request *requests.GetUserRequest) (responses.GetUserResponse, error)
	UpdateUser(request *requests.UpdateUserRequest) (responses.UpdateUserResponse, error)
	DeleteUser(request *requests.DeleteUserRequest) (responses.DeleteUserResponse, error)
	SearchUsers(request *requests.SearchUsersRequest) (responses.SearchUsersResponse, error)
	ChangePassword(request *requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error)

	CreateUserPost(request *requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error)
}

type service struct {
	repository repository.RepositoryLayer
	config     *config.Config
	auth       auth.AuthI
}

func New(repository repository.RepositoryLayer, auth auth.AuthI, config *config.Config) *service {
	return &service{
		repository: repository,
		config:     config,
		auth:       auth,
	}
}
