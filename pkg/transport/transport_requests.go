package transport

import (
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"
	"github.com/gilperopiola/go-rest-example/pkg/requests"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

func makeSignupRequest(c *gin.Context) (request requests.SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.SignupRequest{}, utils.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.SignupRequest{}, err
	}

	return request, nil
}

func makeLoginRequest(c *gin.Context) (request requests.LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.LoginRequest{}, utils.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.LoginRequest{}, err
	}

	return request, nil
}

func makeCreateUserRequest(c *gin.Context) (request requests.CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.CreateUserRequest{}, utils.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.CreateUserRequest{}, err
	}

	return request, nil
}

func makeGetUserRequest(c *gin.Context) (request requests.GetUserRequest, err error) {
	userToGetID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return requests.GetUserRequest{}, err
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return requests.GetUserRequest{}, err
	}

	return request, nil
}

func makeUpdateUserRequest(c *gin.Context) (request requests.UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.UpdateUserRequest{}, utils.Wrap(err, customErrors.ErrBindingRequest)
	}

	userToUpdateID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return requests.UpdateUserRequest{}, err
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return requests.UpdateUserRequest{}, err
	}

	return request, nil
}

func makeDeleteUserRequest(c *gin.Context) (request requests.DeleteUserRequest, err error) {
	userToDeleteID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return requests.DeleteUserRequest{}, err
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return requests.DeleteUserRequest{}, err
	}

	return request, nil
}
