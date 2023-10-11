package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
)

func (s *Service) Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error) {
	user := handlers.New(signupRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return responses.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	user.HashPassword()

	if err := user.Create(s.Repository); err != nil {
		return responses.SignupResponse{}, err
	}

	return responses.SignupResponse{User: user.ToEntity()}, nil
}

func (s *Service) Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error) {
	user := handlers.New(loginRequest.ToUserModel())

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.LoginResponse{}, err
	}

	if !user.PasswordMatches(loginRequest.Password) {
		return responses.LoginResponse{}, entities.ErrWrongPassword
	}

	tokenString, err := user.GenerateTokenString(s.Auth)
	if err != nil {
		return responses.LoginResponse{}, entities.ErrUnauthorized
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
