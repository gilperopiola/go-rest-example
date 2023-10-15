package transport

import (
	"fmt"
	"strconv"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/errors"

	"github.com/gin-gonic/gin"
)

func makeSignupRequest(c *gin.Context) (request requests.SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), err)
	}

	return request, nil
}

func makeLoginRequest(c *gin.Context) (request requests.LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), err)
	}

	return request, nil
}

func makeCreateUserRequest(c *gin.Context) (request requests.CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.CreateUserRequest{}, common.Wrap(fmt.Errorf("makeCreateUserRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return requests.CreateUserRequest{}, common.Wrap(fmt.Errorf("makeCreateUserRequest"), err)
	}

	return request, nil
}

func makeGetUserRequest(c *gin.Context) (request requests.GetUserRequest, err error) {
	userToGetID, err := getIntFromContext(c, "ID")
	if err != nil {
		return requests.GetUserRequest{}, common.Wrap(fmt.Errorf("makeGetUserRequest"), err)
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return requests.GetUserRequest{}, common.Wrap(fmt.Errorf("makeGetUserRequest"), err)
	}

	return request, nil
}

func makeUpdateUserRequest(c *gin.Context) (request requests.UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return requests.UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), customErrors.ErrBindingRequest)
	}

	userToUpdateID, err := getIntFromContext(c, "ID")
	if err != nil {
		return requests.UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), err)
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return requests.UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), err)
	}

	return request, nil
}

func makeDeleteUserRequest(c *gin.Context) (request requests.DeleteUserRequest, err error) {
	userToDeleteID, err := getIntFromContext(c, "ID")
	if err != nil {
		return requests.DeleteUserRequest{}, common.Wrap(fmt.Errorf("makeDeleteUserRequest"), err)
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return requests.DeleteUserRequest{}, common.Wrap(fmt.Errorf("makeDeleteUserRequest"), err)
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
