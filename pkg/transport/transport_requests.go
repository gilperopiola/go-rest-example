package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

func makeSignupRequest(c *gin.Context) (entities.SignupRequest, error) {
	// Bind & validate request
	var signupRequest entities.SignupRequest
	if err := c.ShouldBindJSON(&signupRequest); err != nil {
		return entities.SignupRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	if err := signupRequest.Validate(); err != nil {
		return entities.SignupRequest{}, err
	}

	// Return request
	return signupRequest, nil
}

func makeLoginRequest(c *gin.Context) (entities.LoginRequest, error) {
	// Bind request
	var loginRequest entities.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		return entities.LoginRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	// Validate request
	if err := loginRequest.Validate(); err != nil {
		return entities.LoginRequest{}, err
	}

	// Return request
	return loginRequest, nil
}

func makeCreateUserRequest(c *gin.Context) (entities.CreateUserRequest, error) {
	// Bind & validate request
	var createUserRequest entities.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		return entities.CreateUserRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	if err := createUserRequest.Validate(); err != nil {
		return entities.CreateUserRequest{}, err
	}

	// Return request
	return createUserRequest, nil
}

func makeGetUserRequest(c *gin.Context) (entities.GetUserRequest, error) {
	// Get user ID from context
	userToGetID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.GetUserRequest{}, err
	}

	// Create & validate request
	getUserRequest := entities.GetUserRequest{ID: userToGetID}

	if err := getUserRequest.Validate(); err != nil {
		return entities.GetUserRequest{}, err
	}

	// Return request
	return getUserRequest, nil
}

func makeUpdateUserRequest(c *gin.Context) (entities.UpdateUserRequest, error) {
	// Get user ID from context
	userToUpdateID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.UpdateUserRequest{}, err
	}

	// Bind request
	var updateUserRequest entities.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		return entities.UpdateUserRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	// Assign User ID to request
	updateUserRequest.ID = userToUpdateID

	// Validate request
	if err := updateUserRequest.Validate(); err != nil {
		return entities.UpdateUserRequest{}, err
	}

	// Return request
	return updateUserRequest, nil
}

func makeDeleteUserRequest(c *gin.Context) (entities.DeleteUserRequest, error) {
	// Get user ID from context
	userToDeleteID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.DeleteUserRequest{}, err
	}

	// Create & validate request
	deleteUserRequest := entities.DeleteUserRequest{ID: userToDeleteID}

	if err := deleteUserRequest.Validate(); err != nil {
		return entities.DeleteUserRequest{}, err
	}

	// Return request
	return deleteUserRequest, nil
}
