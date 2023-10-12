package handlers

import (
	"testing"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
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
	testEntity = entities.User{
		ID:       1,
		Username: "username",
		Email:    "email@email.com",
		Password: "password",
	}
)

func TestToEntity(t *testing.T) {
	expected := testEntity
	expected.Password = ""
	got := New(testModel).ToEntity()
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
	mockRepo := repository.NewRepositoryMock()
	mockRepo.On("UserExists", mock.Anything, mock.Anything).Return(true).Once()
	exists := h.Exists(mockRepo)
	assert.True(t, exists)
	mockRepo.AssertExpectations(t)
}

func TestGetAuthRole(t *testing.T) {
	h := New(testModel)
	assert.Equal(t, entities.UserRole, h.GetAuthRole())
	h.User.IsAdmin = true
	assert.Equal(t, entities.AdminRole, h.GetAuthRole())
}

func TestGenerateTokenString(t *testing.T) {
	h := New(testModel)
	mockAuth := new(mockAuth)
	mockAuth.On("GenerateToken", mock.Anything, entities.UserRole).Return("testToken", nil).Once()

	token, err := h.GenerateTokenString(mockAuth)
	assert.NoError(t, err)
	assert.Equal(t, "testToken", token)
	mockAuth.AssertExpectations(t)
}

func TestHashPassword(t *testing.T) {
	h := New(testModel)
	h.HashPassword()
	expected := utils.Hash(testModel.Email, testModel.Password)
	assert.Equal(t, expected, h.User.Password)
}

func TestPasswordMatches(t *testing.T) {
	h := New(testModel)
	h.HashPassword()
	assert.True(t, h.PasswordMatches("password"))
	assert.False(t, h.PasswordMatches("wrong_password"))
}

func TestOverwriteFields(t *testing.T) {
	h := New(testModel)
	h.OverwriteFields("new_username", "new_email@email.com", "new_password")
	assert.Equal(t, "new_username", h.User.Username)
	assert.Equal(t, "new_email@email.com", h.User.Email)
	assert.Equal(t, "new_password", h.User.Password)
}

// -

func getMockRepoWithFnCall(fnName string, userToReturn models.User) *repository.RepositoryMock {
	mockRepo := repository.NewRepositoryMock()
	mockRepo.On(fnName, mock.Anything).Return(userToReturn, nil).Once()
	return mockRepo
}

func assertUserTC(t *testing.T, mock *repository.RepositoryMock, model1 models.User, model2 models.User, err error) {
	assert.NoError(t, err)
	assert.Equal(t, model1, model2)
	mock.AssertExpectations(t)
}
