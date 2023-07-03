package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type Service interface {
	Signup(signupRequest entities.SignupRequest) error
	Login(loginRequest entities.LoginRequest) error
}

type ServiceHandler struct {
	Database *repository.Database
}

func (s *ServiceHandler) Signup(signupRequest entities.SignupRequest) error {
	hashedPassword := utils.Hash(signupRequest.Email, signupRequest.Password)

	user := &models.User{
		Email:    signupRequest.Email,
		Password: hashedPassword,
	}

	if err := s.Database.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}
func (s ServiceHandler) Login(loginRequest entities.LoginRequest) error {
	return nil
}
