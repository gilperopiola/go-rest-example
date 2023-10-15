package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type Service struct {
	Config     config.ConfigI
	Auth       auth.AuthI
	Repository repository.RepositoryLayer
}

type ServiceLayer interface {
	Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error)
	Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error)

	CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error)
	GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error)
	UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error)
	DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error)
}

func NewService(repository repository.RepositoryLayer, auth auth.AuthI, config config.ConfigI) *Service {
	return &Service{
		Repository: repository,
		Auth:       auth,
		Config:     config,
	}
}
