package service

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	VALID_ID         = 1
	VALID_USERNAME   = "valid_username"
	VALID_EMAIL      = "test@email.com"
	VALID_PASSWORD   = "password"
	INVALID_PASSWORD = "invalid_password"
)

func newTestService(mockRepository *repository.RepositoryMock) *Service {
	codec := codec.NewCodec()
	config := &config.Config{}
	auth := &auth.Auth{}
	errorsMapper := ErrorsMapper{}
	return NewService(mockRepository, auth, codec, config, errorsMapper)
}

func TestSignup(t *testing.T) {

	makeMockRepositoryWithUserExists := func(exists bool) *repository.RepositoryMock {
		mockRepository := repository.NewRepositoryMock()
		mockRepository.On("UserExists", mock.Anything, mock.Anything).Return(exists).Once()
		return mockRepository
	}

	makeMockRepositoryWithCreateUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := makeMockRepositoryWithUserExists(false)
		mockRepository.On("CreateUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		request        entities.SignupRequest
		mockRepository *repository.RepositoryMock
		want           entities.SignupResponse
		wantErr        error
	}{
		{
			name:           "error_user_already_exists",
			mockRepository: makeMockRepositoryWithUserExists(true),
			wantErr:        entities.ErrUsernameOrEmailAlreadyInUse,
		},
		{
			name:           "error_creating_user",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, repository.ErrCreatingUser),
			wantErr:        entities.ErrCreatingUser,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithCreateUser(models.User{}, nil),
			want:           entities.SignupResponse{},
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			service := newTestService(tt.mockRepository)

			// Act
			got, err := service.Signup(tt.request)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
			tt.mockRepository.AssertExpectations(t)
		})
	}
}

func TestLogin(t *testing.T) {

	var (
		hashedCorrectPassword = utils.Hash(VALID_EMAIL, VALID_PASSWORD)
		validUser             = models.User{Email: VALID_EMAIL, Password: hashedCorrectPassword}
	)

	makeMockRepositoryWithGetUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := repository.NewRepositoryMock()
		mockRepository.On("GetUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name            string
		request         entities.LoginRequest
		mockRepository  *repository.RepositoryMock
		wantTokenLength int
		wantErr         error
	}{
		{
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, repository.ErrUserNotFound),
			wantErr:        entities.ErrUserNotFound,
		},
		{
			name:           "error_mismatched_passwords",
			mockRepository: makeMockRepositoryWithGetUser(validUser, nil),
			request:        entities.LoginRequest{Password: INVALID_PASSWORD},
			wantErr:        entities.ErrWrongPassword,
		},
		{
			name:            "success",
			mockRepository:  makeMockRepositoryWithGetUser(validUser, nil),
			request:         entities.LoginRequest{Password: VALID_PASSWORD},
			wantErr:         nil,
			wantTokenLength: 212,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			service := newTestService(tt.mockRepository)

			// Act
			got, err := service.Login(tt.request)

			// Assert
			assert.Equal(t, tt.wantTokenLength, len(got.Token))
			assert.ErrorIs(t, err, tt.wantErr)
			tt.mockRepository.AssertExpectations(t)
		})
	}
}

func TestGetUser(t *testing.T) {

	var (
		modelUser  = models.User{Email: VALID_EMAIL}
		entityUser = entities.User{Email: VALID_EMAIL}
	)

	makeMockRepositoryWithGetUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := repository.NewRepositoryMock()
		mockRepository.On("GetUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		request        entities.GetUserRequest
		mockRepository *repository.RepositoryMock
		want           entities.GetUserResponse
		wantErr        error
	}{
		{
			name:           "error_getting_user",
			mockRepository: makeMockRepositoryWithGetUser(models.User{}, repository.ErrUserNotFound),
			wantErr:        entities.ErrUserNotFound,
		},
		{
			name:           "success",
			mockRepository: makeMockRepositoryWithGetUser(modelUser, nil),
			want:           entities.GetUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			service := newTestService(tt.mockRepository)

			// Act
			got, err := service.GetUser(tt.request)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
			tt.mockRepository.AssertExpectations(t)
		})
	}
}

func TestDeleteUser(t *testing.T) {

	var (
		modelUser  = models.User{ID: VALID_ID}
		entityUser = entities.User{ID: VALID_ID}
	)

	makeMockRepositoryWithDeleteUser := func(returnUser models.User, returnErr error) *repository.RepositoryMock {
		mockRepository := repository.NewRepositoryMock()
		mockRepository.On("DeleteUser", mock.Anything).Return(returnUser, returnErr).Once()
		return mockRepository
	}

	tests := []struct {
		name           string
		request        entities.DeleteUserRequest
		mockRepository *repository.RepositoryMock
		want           entities.DeleteUserResponse
		wantErr        error
	}{
		{
			name:           "error_deleting_user",
			request:        entities.DeleteUserRequest{ID: VALID_ID},
			mockRepository: makeMockRepositoryWithDeleteUser(models.User{}, repository.ErrUserAlreadyDeleted),
			wantErr:        entities.ErrUserNotFound,
		},
		{
			name:           "success",
			request:        entities.DeleteUserRequest{ID: VALID_ID},
			mockRepository: makeMockRepositoryWithDeleteUser(modelUser, nil),
			want:           entities.DeleteUserResponse{User: entityUser},
			wantErr:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			service := newTestService(tt.mockRepository)

			// Act
			got, err := service.DeleteUser(tt.request)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
			tt.mockRepository.AssertExpectations(t)
		})
	}
}
