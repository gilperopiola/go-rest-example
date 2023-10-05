package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type Service struct {
	Config       config.ConfigInterface
	Auth         auth.AuthInterface
	Codec        codec.CodecInterface
	Repository   repository.RepositoryLayer
	ErrorsMapper errorsMapperInterface
}

type ServiceLayer interface {
	Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error)
	Login(loginRequest entities.LoginRequest) (entities.LoginResponse, error)

	CreateUser(createUserRequest entities.CreateUserRequest) (entities.CreateUserResponse, error)
	GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error)
	UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error)
	DeleteUser(deleteUserRequest entities.DeleteUserRequest) (entities.DeleteUserResponse, error)
}

func NewService(repository repository.RepositoryLayer, auth auth.AuthInterface, codec codec.CodecInterface, config config.ConfigInterface, errorsMapper errorsMapperInterface) *Service {
	return &Service{
		Repository:   repository,
		Auth:         auth,
		Codec:        codec,
		Config:       config,
		ErrorsMapper: errorsMapper,
	}
}
