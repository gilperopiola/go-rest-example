package service

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
	"github.com/gilperopiola/go-rest-example/pkg/common/mocks"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestService(mockRepository *mocks.RepositoryMock) *service {
	return New(mockRepository, &auth.Auth{}, &config.Config{})
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
	entityUser = responses.User{
		ID:    VALID_ID,
		Email: VALID_EMAIL,
	}
)

func TestSignup(t *testing.T) {
	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("CreateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *mocks.RepositoryMock
		want           responses.SignupResponse
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
			want:           responses.SignupResponse{},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).Signup(requests.SignupRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestLogin(t *testing.T) {

	// The hash here should match the one on the Configuration
	modelUser.Password = common.Hash(VALID_PASSWORD, "")

	tests := []struct {
		name            string
		request         requests.LoginRequest
		mockRepository  *mocks.RepositoryMock
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
			request:        requests.LoginRequest{Password: INVALID_PASSWORD},
			wantErr:        customErrors.ErrWrongPassword,
		},
		{
			name:            "success",
			mockRepository:  makeMockRepositoryWithGetUser(modelUser, nil),
			request:         requests.LoginRequest{Password: VALID_PASSWORD},
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

	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("CreateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *mocks.RepositoryMock
		want           responses.CreateUserResponse
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
			want:           responses.CreateUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).CreateUser(requests.CreateUserRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestGetUser(t *testing.T) {

	tests := []struct {
		name           string
		request        requests.GetUserRequest
		mockRepository *mocks.RepositoryMock
		want           responses.GetUserResponse
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
			want:           responses.GetUserResponse{User: entityUser},
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

	makeMockRepositoryWithGetUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("GetUser", mock.Anything, mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	makeMockRepositoryWithUpdateUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepositoryWithGetUser(returnUser, nil)
		mockRepository.On("UpdateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *mocks.RepositoryMock
		want           responses.UpdateUserResponse
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
			want:           responses.UpdateUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).UpdateUser(requests.UpdateUserRequest{})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		mockRepository *mocks.RepositoryMock
		want           responses.DeleteUserResponse
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
			want:           responses.DeleteUserResponse{User: entityUser},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).DeleteUser(requests.DeleteUserRequest{ID: VALID_ID})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func assertTC(t *testing.T, want interface{}, wantErr error, got interface{}, err error, mockRepository *mocks.RepositoryMock) {
	assert.Equal(t, want, got)
	assert.ErrorIs(t, err, wantErr)
	mockRepository.AssertExpectations(t)
}

func makeMockRepositoryWithUserExists(exists bool) *mocks.RepositoryMock {
	mockRepository := mocks.NewRepositoryMock()
	mockRepository.On("UserExists", mock.Anything, mock.Anything).Return(exists).Once()
	return mockRepository
}

func makeMockRepositoryWithGetUser(returnUser models.User, returnErr error) *mocks.RepositoryMock {
	mockRepository := mocks.NewRepositoryMock()
	mockRepository.On("GetUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}

func makeMockRepositoryWithDeleteUser(returnUser models.User, returnErr error) *mocks.RepositoryMock {
	mockRepository := mocks.NewRepositoryMock()
	mockRepository.On("DeleteUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}
