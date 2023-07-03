package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type Service interface {
	Signup(signupRequest entities.SignupRequest) error
	Login(loginRequest entities.LoginRequest) error
}

type ServiceHandler struct {
	Database   *repository.Database
	Repository *repository.RepositoryHandler
	Codec      *codec.CodecHandler
}

func (s *ServiceHandler) Signup(signupRequest entities.SignupRequest) error {

	// Validations
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username) {
		return entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Actions
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transformations
	user := s.Codec.FromSignupRequestToUserModel(signupRequest, hashedPassword)

	// Database
	if err := s.Repository.CreateUser(&user); err != nil {
		return err
	}

	return nil
}
func (s ServiceHandler) Login(loginRequest entities.LoginRequest) error {
	return nil
}
