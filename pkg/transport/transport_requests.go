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
		return entities.SignupRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
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
		return entities.LoginRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
	}

	// Validate request
	if err := loginRequest.Validate(); err != nil {
		return entities.LoginRequest{}, err
	}

	// Return request
	return loginRequest, nil
}

func makeCreateUserRequest(c *gin.Context) (entities.CreateUserRequest, error) {
	return entities.CreateUserRequest{}, nil
}
func makeGetUserRequest(c *gin.Context) (entities.GetUserRequest, error) {

	// Get info from context and URL, check if user IDs match
	userToGetID, err := getUserIDFromContext(c)
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

	// Validate and get User ID
	userToUpdateID, err := getUserIDFromContext(c)
	if err != nil {
		return entities.UpdateUserRequest{}, err
	}

	// Bind request
	var updateUserRequest entities.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		return entities.UpdateUserRequest{}, utils.JoinErrors(entities.ErrBindingRequest, err)
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

	// Get info from context and URL, check if user IDs match
	userToDeleteID, err := getUserIDFromContext(c)
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

/* ----------------- */

func getUserIDFromContext(c *gin.Context) (int, error) {

	// Get logged user ID
	loggedUserID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return 0, err
	}

	// Get URL user ID
	userToGetID, err := utils.GetIntFromContextParams(c.Params, "user_id")
	if err != nil {
		return 0, err
	}

	// Check if the logged user has the same ID as the one to get
	if loggedUserID != userToGetID {
		return 0, entities.ErrUnauthorized
	}

	// Return user ID
	return userToGetID, nil
}
