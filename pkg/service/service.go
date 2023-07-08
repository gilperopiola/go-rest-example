package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type Service struct {
	Repository   repository.RepositoryIFace
	Codec        codec.CodecIFace
	Config       config.Config
	ErrorsMapper ErrorsMapperIface
}

type ServiceIFace interface {
	Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error)
	Login(userCredentials entities.UserCredentials) (entities.LoginResponse, error)
	GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error)
	UpdateUser(updateUserRequest entities.UpdateUserRequest) (entities.UpdateUserResponse, error)
}

func NewService(repository repository.RepositoryIFace, codec codec.CodecIFace, config config.Config, errorsMapper ErrorsMapperIface) *Service {
	return &Service{
		Repository:   repository,
		Codec:        codec,
		Config:       config,
		ErrorsMapper: errorsMapper,
	}
}
