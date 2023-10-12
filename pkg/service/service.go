package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type Service struct {
	Config     config.ConfigI
	Auth       auth.AuthI
	Repository repository.RepositoryLayer
}

type ServiceLayer interface {
	Signup(signupRequest common.SignupRequest) (common.SignupResponse, error)
	Login(loginRequest common.LoginRequest) (common.LoginResponse, error)

	CreateUser(createUserRequest common.CreateUserRequest) (common.CreateUserResponse, error)
	GetUser(getUserRequest common.GetUserRequest) (common.GetUserResponse, error)
	UpdateUser(updateUserRequest common.UpdateUserRequest) (common.UpdateUserResponse, error)
	DeleteUser(deleteUserRequest common.DeleteUserRequest) (common.DeleteUserResponse, error)
}

func NewService(repository repository.RepositoryLayer, auth auth.AuthI, config config.ConfigI) *Service {
	return &Service{
		Repository: repository,
		Auth:       auth,
		Config:     config,
	}
}
