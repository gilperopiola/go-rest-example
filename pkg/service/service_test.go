package service

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/codec"
	"github.com/gilperopiola/go-rest-example/pkg/config"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignup(t *testing.T) {

	tests := []struct {
		name           string
		request        entities.SignupRequest
		mockRepository func() *repository.RepositoryMock
		want           entities.SignupResponse
		wantErr        error
	}{
		{
			name: "error_user_already_exists",
			mockRepository: func() *repository.RepositoryMock {
				mockRepository := &repository.RepositoryMock{Mock: &mock.Mock{}}
				mockRepository.On("UserExists", mock.Anything, mock.Anything).Return(true).Once()
				return mockRepository
			},
			wantErr: entities.ErrUsernameOrEmailAlreadyInUse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			codec := &codec.Codec{}
			config := config.Config{}
			errorsMapper := ErrorsMapper{}

			service := NewService(tt.mockRepository(), codec, config, errorsMapper)

			// Act
			got, err := service.Signup(tt.request)

			// Assert
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}
