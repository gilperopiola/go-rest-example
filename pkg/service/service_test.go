package service

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestService(mockRepository *repository.RepositoryMock) *Service {
	return NewService(mockRepository, &auth.Auth{}, &config.Config{})
}

const (
	VALID_ID         = 1
	VALID_USERNAME   = "valid_username"
	VALID_EMAIL      = "test@email.com"
	VALID_PASSWORD   = "password"
	INVALID_PASSWORD = "invalid_password"
)

var (
	modelUser = models.User{
		ID:    VALID_ID,
		Email: VALID_EMAIL,
	}
	entityUser = entities.User{
		ID:    VALID_ID,
		Email: VALID_EMAIL,
	}
)

func TestSignup(t *testing.T) {
	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("CreateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *repository.RepositoryMock
		want           common.SignupResponse
		wantErr        error
	}{
		{
			name:           "error_user_already_exists",
			mockRepository: makeMockRepositoryWithUserExists(true),
			wantErr:        customErrors.ErrUsernameOrEmailAlreadyInUse,
		},
		{
			name:           "error_creating_user",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, customErrors.ErrCreatingUser),
			wantErr:        customErrors.ErrCreatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, nil),
			want:           common.SignupResponse{},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).Signup(common.SignupRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestLogin(t *testing.T) {

	modelUser.Password = common.Hash(VALID_EMAIL, VALID_PASSWORD)

	tests := []struct {
		name            string
		request         common.LoginRequest
		mockRepository  *repository.RepositoryMock
		wantTokenLength int
		wantErr         error
	}{
		{
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, customErrors.ErrUserNotFound),
			wantErr:        customErrors.ErrUserNotFound,
		},
		{
			name:           "error_mismatched_passwords",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, nil),
			request:        common.LoginRequest{Password: INVALID_PASSWORD},
			wantErr:        customErrors.ErrWrongPassword,
		},
		{
			name:            "success",
			mockRepository:  makeMockRepositoryWithGetUser(modelUser, nil),
			request:         common.LoginRequest{Password: VALID_PASSWORD},
			wantErr:         nil,
			wantTokenLength: 212,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).Login(tc.request)
			assertTC(t, tc.wantTokenLength, tc.wantErr, len(got.Token), err, tc.mockRepository)
		})
	}
}

func TestCreateUser(t *testing.T) {

	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("CreateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *repository.RepositoryMock
		want           common.CreateUserResponse
		wantErr        error
	}{
		{
			name:           "error_user_exists",
			mockRepository: makeMockRepositoryWithUserExists(true),
			wantErr:        customErrors.ErrUsernameOrEmailAlreadyInUse,
		},
		{
			name:           "error_creating_user",
			mockRepository: makeMockRepositoryWithCreateUser(modelUser, customErrors.ErrCreatingUser),
			wantErr:        customErrors.ErrCreatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithCreateUser(modelUser, nil),
			want:           common.CreateUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).CreateUser(common.CreateUserRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestGetUser(t *testing.T) {

	tests := []struct {
		name           string
		request        common.GetUserRequest
		mockRepository *repository.RepositoryMock
		want           common.GetUserResponse
		wantErr        error
	}{
		{
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, customErrors.ErrUserNotFound),
			wantErr:        customErrors.ErrUserNotFound,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, nil),
			want:           common.GetUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).GetUser(tc.request)
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestUpdateUser(t *testing.T) {

	makeMockRepositoryWithGetUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("GetUser", mock.Anything, mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	makeMockRepositoryWithUpdateUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := makeMockRepositoryWithGetUser(returnUser, nil)
		mockRepository.On("UpdateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *repository.RepositoryMock
		want           common.UpdateUserResponse
		wantErr        error
	}{
		{
			name:           "error_user_exists",
			mockRepository: makeMockRepositoryWithUserExists(true),
			wantErr:        customErrors.ErrUsernameOrEmailAlreadyInUse,
		},
		{
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, customErrors.ErrGettingUser),
			wantErr:        customErrors.ErrGettingUser,
		},
		{
			name:           "error_updating_user",
			mockRepository: makeMockRepositoryWithUpdateUser(modelUser, customErrors.ErrUpdatingUser),
			wantErr:        customErrors.ErrUpdatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithUpdateUser(modelUser, nil),
			want:           common.UpdateUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).UpdateUser(common.UpdateUserRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		mockRepository *repository.RepositoryMock
		want           common.DeleteUserResponse
		wantErr        error
	}{
		{
			name:           "error_deleting_user",
			mockRepository: makeMockRepositoryWithDeleteUser(models.User{}, customErrors.ErrUserAlreadyDeleted),
			wantErr:        customErrors.ErrUserAlreadyDeleted,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithDeleteUser(modelUser, nil),
			want:           common.DeleteUserResponse{User: entityUser},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).DeleteUser(common.DeleteUserRequest{ID: VALID_ID})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func assertTC(t *testing.T, want interface{}, wantErr error, got interface{}, err error, mockRepository *repository.RepositoryMock) {
	assert.Equal(t, want, got)
	assert.ErrorIs(t, err, wantErr)
	mockRepository.AssertExpectations(t)
}

func makeMockRepositoryWithUserExists(exists bool) *repository.RepositoryMock {
	mockRepository := repository.NewRepositoryMock()
	mockRepository.On("UserExists", mock.Anything, mock.Anything).Return(exists).Once()
	return mockRepository
}

func makeMockRepositoryWithGetUser(returnUser models.User, returnErr error) *repository.RepositoryMock {
	mockRepository := repository.NewRepositoryMock()
	mockRepository.On("GetUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}

func makeMockRepositoryWithDeleteUser(returnUser models.User, returnErr error) *repository.RepositoryMock {
	mockRepository := repository.NewRepositoryMock()
	mockRepository.On("DeleteUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}
