package handlers

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/mocks"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	testModel = models.User{
		ID:       1,
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	testResponseModel = responses.User{
		ID:       1,
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
	testAuthEntity = auth.User{
		ID:       1,
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
)

func TestToResponseModel(t *testing.T) {
	expected := testResponseModel
	expected.Password = ""
	got := New(testModel).ToResponseModel()
	assert.Equal(t, expected, got)
}

func TestToAuthEntity(t *testing.T) {
	expected := testAuthEntity
	expected.Password = ""
	got := New(testModel).ToAuthEntity()
	assert.Equal(t, expected, got)
}

func TestCreate(t *testing.T) {
	h := New(testModel)
	mockRepo := getMockRepoWithFnCall("CreateUser", testModel)
	err := h.Create(mockRepo)
	assertUserTC(t, mockRepo, testModel, h.User, err)
}

func TestGet(t *testing.T) {
	h := New(testModel)
	mockRepo := getMockRepoWithFnCall("GetUser", testModel)
	err := h.Get(mockRepo)
	assertUserTC(t, mockRepo, testModel, h.User, err)
}

func TestUpdate(t *testing.T) {
	h := New(testModel)
	mockRepo := getMockRepoWithFnCall("UpdateUser", testModel)
	err := h.Update(mockRepo)
	assertUserTC(t, mockRepo, testModel, h.User, err)
}

func TestDelete(t *testing.T) {
	h := New(testModel)
	mockRepo := getMockRepoWithFnCall("DeleteUser", testModel)
	err := h.Delete(mockRepo)
	assertUserTC(t, mockRepo, testModel, h.User, err)
}

func TestExists(t *testing.T) {
	h := New(testModel)
	mockRepo := mocks.NewRepositoryMock()
	mockRepo.On("UserExists", mock.Anything, mock.Anything).Return(true).Once()
	exists := h.Exists(mockRepo)
	assert.True(t, exists)
	mockRepo.AssertExpectations(t)
}

func TestGetAuthRole(t *testing.T) {
	h := New(testModel)
	assert.Equal(t, auth.UserRole, h.GetAuthRole())
	h.User.IsAdmin = true
	assert.Equal(t, auth.AdminRole, h.GetAuthRole())
}

func TestGenerateTokenString(t *testing.T) {
	h := New(testModel)
	mockAuth := new(mocks.MockAuth)
	mockAuth.On("GenerateToken", mock.Anything, auth.UserRole).Return("testToken", nil).Once()

	token, err := h.GenerateTokenString(mockAuth)
	assert.NoError(t, err)
	assert.Equal(t, "testToken", token)
	mockAuth.AssertExpectations(t)
}

func TestHashPassword(t *testing.T) {
	h := New(testModel)
	h.HashPassword("salt")
	expected := common.Hash(testModel.Password, "salt")
	assert.Equal(t, expected, h.User.Password)
}

func TestPasswordMatches(t *testing.T) {
	h := New(testModel)
	h.HashPassword("salt")
	assert.True(t, h.PasswordMatches("password", "salt"))
	assert.False(t, h.PasswordMatches("wrong_password", "salt"))
}

func TestOverwriteFields(t *testing.T) {
	h := New(testModel)
	h.OverwriteFields("new_username", "new_email@email.com", "new_password")
	assert.Equal(t, "new_username", h.User.Username)
	assert.Equal(t, "new_email@email.com", h.User.Email)
	assert.Equal(t, "new_password", h.User.Password)
}

// - Helpers

func getMockRepoWithFnCall(fnName string, userToReturn models.User) *mocks.RepositoryMock {
	mockRepo := mocks.NewRepositoryMock()
	mockRepo.On(fnName, mock.Anything).Return(userToReturn, nil).Once()
	return mockRepo
}

func assertUserTC(t *testing.T, mock *mocks.RepositoryMock, model1 models.User, model2 models.User, err error) {
	assert.NoError(t, err)
	assert.Equal(t, model1, model2)
	mock.AssertExpectations(t)
}
