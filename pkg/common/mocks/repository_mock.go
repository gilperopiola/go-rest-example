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

func (m *RepositoryMock) UpdateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) GetUser(user models.User, opts ...options.QueryOption) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) UserExists(username, email string, opts ...options.QueryOption) bool {
	args := m.Called(username, email)
	return args.Bool(0)
}

func (m *RepositoryMock) DeleteUser(id int) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *RepositoryMock) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	args := m.Called(post)
	return args.Get(0).(models.UserPost), args.Error(1)
}
