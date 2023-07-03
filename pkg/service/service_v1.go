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
	Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error)
	Login(loginRequest entities.LoginRequest) error
}

type Service struct {
	Database   repository.Databaser
	Repository repository.Repositorier
	Codec      codec.Codecer
}

func (s *Service) Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error) {

	// Validate user doesn't exist
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username) {
		return entities.SignupResponse{}, entities.ErrUsernameOrEmailAlreadyInUse
	}

	// Hash password
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transform request to model
	user := s.Codec.FromSignupRequestToUserModel(signupRequest, hashedPassword)

	// Create user model on the database
	createdUser, err := s.Repository.CreateUser(user)
	if err != nil {
		if errors.Is(err, repository.ErrCreatingUser) {
			return entities.SignupResponse{}, fmt.Errorf("%w:%w", entities.ErrCreatingUser, err)
		}
		return entities.SignupResponse{}, err
	}

	// Transform user to entities
	encodedUser := s.Codec.FromUserModelToEntities(createdUser)

	// Return response
	return entities.SignupResponse{User: encodedUser}, nil
}
func (s Service) Login(loginRequest entities.LoginRequest) error {
	return nil
}
