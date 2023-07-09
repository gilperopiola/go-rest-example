package repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/models"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	*mock.Mock
}

func (m *RepositoryMock) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) UpdateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) GetUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) UserExists(email, username string) bool {
	args := m.Called(email, username)
	return args.Bool(0)
}
