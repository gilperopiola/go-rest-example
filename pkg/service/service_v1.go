package service

import (
	"errors"
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type Servicer interface {
	Signup(signupRequest entities.SignupRequest) error
	Login(loginRequest entities.LoginRequest) error
}

type Service struct {
	Database   *repository.Database
	Repository *repository.Repository
	Codec      *codec.Codec
}

func (s *Service) Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error) {

	// Validations
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username) {
		return entities.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Actions
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transformations
	user := s.Codec.FromSignupRequestToUserModel(signupRequest, hashedPassword)

	// Database
	createdUser, err := s.Repository.CreateUser(user)
	if err != nil {
		if errors.Is(err, repository.ErrCreatingUser) {
			return entities.SignupResponse{}, fmt.Errorf("%w:%w", entities.ErrCreatingUser, err)
		}
		return entities.SignupResponse{}, err
	}

	// Transform response
	encodedUser := s.Codec.FromUserModelToEntities(createdUser)

	return entities.SignupResponse{User: encodedUser}, nil
}
func (s Service) Login(loginRequest entities.LoginRequest) error {
	return nil
}
