package service

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
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
	modelUser    = models.User{ID: VALID_ID, Email: VALID_EMAIL, Posts: []models.UserPost{}}
	responseUser = responses.User{ID: VALID_ID, Email: VALID_EMAIL, Posts: []responses.UserPost{}}
)

func TestSignup(t *testing.T) {
	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepository()
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
			name:           "error_creating_user",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, common.ErrCreatingUser),
			wantErr:        common.ErrCreatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, nil),
			want:           responses.SignupResponse{User: responses.User{Posts: []responses.UserPost{}}},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).Signup(&requests.SignupRequest{})
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
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, common.ErrUserNotFound),
			wantErr:        common.ErrUserNotFound,
		},
		{
			name:           "error_mismatched_passwords",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, nil),
			request:        requests.LoginRequest{Password: INVALID_PASSWORD},
			wantErr:        common.ErrWrongPassword,
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
			got, err := newTestService(tc.mockRepository).Login(&tc.request)
			assertTC(t, tc.wantTokenLength, tc.wantErr, len(got.Token), err, tc.mockRepository)
		})
	}
}

func TestCreateUser(t *testing.T) {

	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepository()
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
			name:           "error_creating_user",
			mockRepository: makeMockRepositoryWithCreateUser(modelUser, common.ErrCreatingUser),
			wantErr:        common.ErrCreatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithCreateUser(modelUser, nil),
			want:           responses.CreateUserResponse{User: responseUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).CreateUser(&requests.CreateUserRequest{})
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
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, common.ErrUserNotFound),
			wantErr:        common.ErrUserNotFound,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, nil),
			want:           responses.GetUserResponse{User: responseUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).GetUser(&tc.request)
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestUpdateUser(t *testing.T) {

	makeMockRepositoryWithGetUser := func(returnUser models.User, returnErr error) *mocks.RepositoryMock {
		mockRepository := makeMockRepository()
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
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, common.ErrGettingUser),
			wantErr:        common.ErrGettingUser,
		},
		{
			name:           "error_updating_user",
			mockRepository: makeMockRepositoryWithUpdateUser(modelUser, common.ErrUpdatingUser),
			wantErr:        common.ErrUpdatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithUpdateUser(modelUser, nil),
			want:           responses.UpdateUserResponse{User: responseUser},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).UpdateUser(&requests.UpdateUserRequest{})
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
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, common.ErrUserAlreadyDeleted),
			wantErr:        common.ErrUserAlreadyDeleted,
		},
		{
			name:           "error_deleting_user",
			mockRepository: makeMockRepositoryWithDeleteUser(models.User{}, common.ErrUserAlreadyDeleted),
			wantErr:        common.ErrUserAlreadyDeleted,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithDeleteUser(modelUser, nil),
			want:           responses.DeleteUserResponse{User: responseUser},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).DeleteUser(&requests.DeleteUserRequest{UserID: VALID_ID})
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func TestSearchUsers(t *testing.T) {

	makeMockRepositoryWithSearchUsers := func(returnUsers models.Users, returnErr error) *mocks.RepositoryMock {
		mockRepository := mocks.NewRepositoryMock()
		mockRepository.On("SearchUsers", mock.Anything, mock.Anything, mock.Anything).Return(returnUsers, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		mockRepository *mocks.RepositoryMock
		request        requests.SearchUsersRequest
		want           responses.SearchUsersResponse
		wantErr        error
	}{
		{
			name:           "error_searching_users",
			mockRepository: makeMockRepositoryWithSearchUsers(nil, common.ErrSearchingUsers),
			request:        requests.SearchUsersRequest{Page: 1, PerPage: 10},
			wantErr:        common.ErrSearchingUsers,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithSearchUsers([]models.User{modelUser}, nil),
			request:        requests.SearchUsersRequest{Page: 1, PerPage: 10},
			want:           responses.SearchUsersResponse{Users: []responses.User{responseUser}, Page: 1, PerPage: 10},
			wantErr:        nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := newTestService(tc.mockRepository).SearchUsers(&tc.request)
			assertTC(t, tc.want, tc.wantErr, got, err, tc.mockRepository)
		})
	}
}

func assertTC(t *testing.T, want interface{}, wantErr error, got interface{}, err error, mockRepository *mocks.RepositoryMock) {
	assert.Equal(t, want, got)
	assert.ErrorIs(t, err, wantErr)
	mockRepository.AssertExpectations(t)
}

func makeMockRepository() *mocks.RepositoryMock {
	return mocks.NewRepositoryMock()
}

func makeMockRepositoryWithGetUser(returnUser models.User, returnErr error) *mocks.RepositoryMock {
	mockRepository := mocks.NewRepositoryMock()
	mockRepository.On("GetUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}

func makeMockRepositoryWithDeleteUser(returnUser models.User, returnErr error) *mocks.RepositoryMock {
	mockRepository := makeMockRepositoryWithGetUser(returnUser, nil)
	mockRepository.On("DeleteUser", mock.Anything).Return(returnUser, returnErr).Once()
	return mockRepository
}
