package mocks

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	*mock.Mock
}

func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{Mock: &mock.Mock{}}
}

func (m *RepositoryMock) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) UpdateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) UpdatePassword(userID int, password string) error {
	args := m.Called(userID, password)
	return args.Error(0)
}

func (m *RepositoryMock) DeleteUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) SearchUsers(page, perPage int, opts ...options.QueryOption) (models.Users, error) {
	args := m.Called(page, perPage)
	return args.Get(0).(models.Users), args.Error(1)
}

func (m *RepositoryMock) UserExists(username, email string, opts ...options.QueryOption) bool {
	args := m.Called(username, email)
	return args.Bool(0)
}

func (m *RepositoryMock) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	args := m.Called(post)
	return args.Get(0).(models.UserPost), args.Error(1)
}
