package transport

import (
	"fmt"
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"

	"github.com/gin-gonic/gin"
)

func makeSignupRequest(c *gin.Context) (request common.SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return common.SignupRequest{}, common.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return common.SignupRequest{}, err
	}

	return request, nil
}

func makeLoginRequest(c *gin.Context) (request common.LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return common.LoginRequest{}, common.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return common.LoginRequest{}, err
	}

	return request, nil
}

func makeCreateUserRequest(c *gin.Context) (request common.CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return common.CreateUserRequest{}, common.Wrap(err, customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return common.CreateUserRequest{}, err
	}

	return request, nil
}

func makeGetUserRequest(c *gin.Context) (request common.GetUserRequest, err error) {
	userToGetID, err := getIntFromContext(c, "ID")
	if err != nil {
		return common.GetUserRequest{}, err
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return common.GetUserRequest{}, err
	}

	return request, nil
}

func makeUpdateUserRequest(c *gin.Context) (request common.UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return common.UpdateUserRequest{}, common.Wrap(err, customErrors.ErrBindingRequest)
	}

	userToUpdateID, err := getIntFromContext(c, "ID")
	if err != nil {
		return common.UpdateUserRequest{}, err
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return common.UpdateUserRequest{}, err
	}

	return request, nil
}

func makeDeleteUserRequest(c *gin.Context) (request common.DeleteUserRequest, err error) {
	userToDeleteID, err := getIntFromContext(c, "ID")
	if err != nil {
		return common.DeleteUserRequest{}, err
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return common.DeleteUserRequest{}, err
	}

	return request, nil
}

func getIntFromContext(c *gin.Context, key string) (int, error) {

	// Get from context
	value, ok := c.Get(key)
	if value == nil || !ok {
		return 0, fmt.Errorf("error getting %s from context", key)
	}

	// Cast to string
	valueStr, ok := value.(string)
	if !ok {
		return 0, fmt.Errorf("error casting %s from context to string", key)
	}

	// Convert to int
	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, fmt.Errorf("error converting %s from string to int", key)
	}

	return valueInt, nil
}
