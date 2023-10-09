package transport

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

func makeSignupRequest(c *gin.Context) (request entities.SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return entities.SignupRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return entities.SignupRequest{}, err
	}

	return request, nil
}

func makeLoginRequest(c *gin.Context) (request entities.LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return entities.LoginRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return entities.LoginRequest{}, err
	}

	return request, nil
}

func makeCreateUserRequest(c *gin.Context) (request entities.CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return entities.CreateUserRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return entities.CreateUserRequest{}, err
	}

	return request, nil
}

func makeGetUserRequest(c *gin.Context) (request entities.GetUserRequest, err error) {
	userToGetID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.GetUserRequest{}, err
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return entities.GetUserRequest{}, err
	}

	return request, nil
}

func makeUpdateUserRequest(c *gin.Context) (request entities.UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return entities.UpdateUserRequest{}, utils.Wrap(err, entities.ErrBindingRequest)
	}

	userToUpdateID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.UpdateUserRequest{}, err
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return entities.UpdateUserRequest{}, err
	}

	return request, nil
}

func makeDeleteUserRequest(c *gin.Context) (request entities.DeleteUserRequest, err error) {
	userToDeleteID, err := utils.GetIntFromContext(c, "ID")
	if err != nil {
		return entities.DeleteUserRequest{}, err
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return entities.DeleteUserRequest{}, err
	}

	return request, nil
}
