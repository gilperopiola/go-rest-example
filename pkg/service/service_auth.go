package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
)

func (s *Service) Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error) {
	userHandler := handlers.New(signupRequest.ToUserModel())

	if userHandler.Exists(s.Repository) {
		return responses.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	userHandler.HashPassword()

	if err := userHandler.Create(s.Repository); err != nil {
		return responses.SignupResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.SignupResponse{User: userHandler.ToEntity()}, nil
}

func (s *Service) Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error) {
	userHandler := handlers.New(loginRequest.ToUserModel())

	if err := userHandler.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	if !userHandler.PasswordMatches(loginRequest.Password) {
		return responses.LoginResponse{}, entities.ErrWrongPassword
	}

	tokenString, err := userHandler.GenerateTokenString(s.Auth)
	if err != nil {
		return responses.LoginResponse{}, entities.ErrUnauthorized
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
