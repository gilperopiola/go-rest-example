package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/service/handlers"
)

func (s *Service) Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error) {
	user := handlers.New(signupRequest.ToUserModel())

	if user.Exists(s.Repository) {
		return responses.SignupResponse{}, common.Wrap(fmt.Errorf("Signup: user.Exists"), customErrors.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword()

	if err := user.Create(s.Repository); err != nil {
		return responses.SignupResponse{}, common.Wrap(fmt.Errorf("Signup: user.Create"), err)
	}

	return responses.SignupResponse{User: user.ToEntity()}, nil
}

func (s *Service) Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error) {
	user := handlers.New(loginRequest.ToUserModel())

	if err := user.Get(s.Repository, repository.WithoutDeleted); err != nil {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: user.Get"), err)
	}

	if !user.PasswordMatches(loginRequest.Password) {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: !user.PasswordMatches"), customErrors.ErrWrongPassword)
	}

	tokenString, err := user.GenerateTokenString(s.Auth)
	if err != nil {
		return responses.LoginResponse{}, common.Wrap(fmt.Errorf("Login: user.GenerateTokenString"), customErrors.ErrUnauthorized)
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
