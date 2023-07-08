package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (s *Service) Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error) {

	// Validate user doesn't exist
	if s.Repository.UserExists(signupRequest.Email, signupRequest.Username) {
		return entities.SignupResponse{}, s.ErrorsMapper.Map(entities.ErrUsernameOrEmailAlreadyInUse)
	}

	// Hash password
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	// Transform request to model
	user := s.Codec.FromSignupRequestToUserModel(signupRequest, hashedPassword)

	// Create user model on the database
	createdUser, err := s.Repository.CreateUser(user)
	if err != nil {
		return entities.SignupResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user to entities
	encodedUser := s.Codec.FromUserModelToEntities(createdUser)

	// Return response
	return entities.SignupResponse{User: encodedUser}, nil
}
