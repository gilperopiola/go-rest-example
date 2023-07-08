package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type Service struct {
	Repository   repository.RepositoryIFace
	Codec        codec.CodecIFace
	Config       config.Config
	ErrorsMapper ErrorsMapperIface
}

func NewService(repository repository.RepositoryIFace, codec codec.CodecIFace, config config.Config, errorsMapper ErrorsMapperIface) *Service {
	return &Service{
		Repository:   repository,
		Codec:        codec,
		Config:       config,
		ErrorsMapper: errorsMapper,
	}
}

type ServiceIFace interface {
	Signup(signupRequest entities.SignupRequest) (entities.SignupResponse, error)
	Login(userCredentials entities.UserCredentials) (entities.LoginResponse, error)
	GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error)
}

//----------------------------------------

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

func (s *Service) Login(userCredentials entities.UserCredentials) (entities.LoginResponse, error) {
	// Transform user entity to model
	userModel := s.Codec.FromUserCredentialsToUserModel(userCredentials)

	// Get user from database
	user, err := s.Repository.GetUser(userModel)
	if err != nil {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(err)
	}

	// Check if passwords match
	if user.Password != utils.Hash(user.Email, userCredentials.Password) {
		return entities.LoginResponse{}, s.ErrorsMapper.Map(entities.ErrWrongPassword)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(user)

	// Generate and return token
	return entities.LoginResponse{
		Token: auth.GenerateToken(userEntity, s.Config.JWT.SESSION_DURATION_DAYS, s.Config.JWT.SECRET),
	}, nil
}

func (s *Service) GetUser(getUserRequest entities.GetUserRequest) (entities.GetUserResponse, error) {

	// Create user model for DB searching
	userModel := models.User{
		ID: getUserRequest.ID,
	}

	// Get user from database
	user, err := s.Repository.GetUser(userModel)
	if err != nil {
		return entities.GetUserResponse{}, s.ErrorsMapper.Map(err)
	}

	// Transform user model to entity
	userEntity := s.Codec.FromUserModelToEntities(user)

	// Return user
	return entities.GetUserResponse{User: userEntity}, nil
}
