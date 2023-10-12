package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

func (s *Service) Signup(signupRequest common.SignupRequest) (common.SignupResponse, error) {
	user := handlers.New(signupRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return common.SignupResponse{}, customErrors.ErrUsernameOrEmailAlreadyInUse
	}

	user.HashPassword()

	if err := user.Create(s.Repository); err != nil {
		return common.SignupResponse{}, err
	}

	return common.SignupResponse{User: user.ToEntity()}, nil
}

func (s *Service) Login(loginRequest common.LoginRequest) (common.LoginResponse, error) {
	user := handlers.New(loginRequest.ToUserModel())

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return common.LoginResponse{}, err
	}

	if !user.PasswordMatches(loginRequest.Password) {
		return common.LoginResponse{}, customErrors.ErrWrongPassword
	}

	tokenString, err := user.GenerateTokenString(s.Auth)
	if err != nil {
		return common.LoginResponse{}, customErrors.ErrUnauthorized
	}

	return common.LoginResponse{Token: tokenString}, nil
}
