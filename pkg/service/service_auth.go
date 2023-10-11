package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/handlers"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/responses"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Signup(signupRequest requests.SignupRequest) (responses.SignupResponse, error) {
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username) {
		return responses.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	signupRequest.Password = utils.Hash(signupRequest.Email, signupRequest.Password)

	userHandler := handlers.New(signupRequest.ToUserModel())

	if err := userHandler.Create(s.Repository); err != nil {
		return responses.SignupResponse{}, s.ErrorsMapper.Map(err)
	}

	return responses.SignupResponse{User: userHandler.User.ToEntity()}, nil
}

func (s *Service) Login(loginRequest requests.LoginRequest) (responses.LoginResponse, error) {
	userFromDB, err := s.Repository.GetUser(loginRequest.ToUserModel(), repository.WithoutDeleted)
	if err != nil {
		return responses.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	if !userFromDB.PasswordMatches(loginRequest.Password) {
		return responses.LoginResponse{}, entities.ErrWrongPassword
	}

	tokenString, err := s.Auth.GenerateToken(userFromDB.ToEntity(), userFromDB.GetAuthRole())
	if err != nil {
		return responses.LoginResponse{}, entities.ErrUnauthorized
	}

	return responses.LoginResponse{Token: tokenString}, nil
}
