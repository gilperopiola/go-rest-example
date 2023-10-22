package repository

import (
	"fmt"
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/mocks"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	errGeneric = fmt.Errorf("error")
)

func TestCreateUser(t *testing.T) {
	makeDBMockWithCreate := func(errToReturn error) *mocks.DBMock {
		mockDB := &mocks.DBMock{}
		mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: errToReturn}).Once()
		return mockDB
	}

	tests := []struct {
		name    string
		dbMock  *mocks.DBMock
		wantErr error
	}{
		{
			name:    "error_creating_user",
			dbMock:  makeDBMockWithCreate(errGeneric),
			wantErr: common.ErrCreatingUser,
		},
		{
			name:   "success",
			dbMock: makeDBMockWithCreate(nil),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// prepare
			repo := New(database{db: tc.dbMock})

			// act
			_, err := repo.CreateUser(models.User{})

			// assert
			assert.ErrorIs(t, err, tc.wantErr)
			tc.dbMock.AssertExpectations(t)
		})
	}
}
