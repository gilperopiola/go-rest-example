package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type Service struct {
	Config       config.ConfigProvider
	Auth         auth.AuthProvider
	Codec        codec.CodecProvider
	Repository   repository.RepositoryProvider
	ErrorsMapper ErrorsMapperProvider
}

type ServiceProvider interface {
	Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error)
	Login(loginRequest entities.LoginRequest) (entities.LoginResponse, error)

	GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error)
	UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error)
	DeleteUser(deleteUserRequest entities.DeleteUserRequest) (entities.DeleteUserResponse, error)
}

func NewService(repository repository.RepositoryProvider, auth auth.AuthProvider, codec codec.CodecProvider, config config.ConfigProvider, errorsMapper ErrorsMapperProvider) *Service {
	return &Service{
		Repository:   repository,
		Auth:         auth,
		Codec:        codec,
		Config:       config,
		ErrorsMapper: errorsMapper,
	}
}
