package requests

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	customErrors "github.com/gilperopiola/go-rest-example/pkg/common/errors"
)

type GinI interface {
	ShouldBindJSON(obj interface{}) error
	GetInt(key string) int
}

func MakeSignupRequest(c GinI) (request SignupRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return SignupRequest{}, common.Wrap(fmt.Errorf("makeSignupRequest"), err)
	}

	return request, nil
}

func MakeLoginRequest(c GinI) (request LoginRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return LoginRequest{}, common.Wrap(fmt.Errorf("makeLoginRequest"), err)
	}

	return request, nil
}

func MakeCreateUserRequest(c GinI) (request CreateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return CreateUserRequest{}, common.Wrap(fmt.Errorf("makeCreateUserRequest"), customErrors.ErrBindingRequest)
	}

	if err = request.Validate(); err != nil {
		return CreateUserRequest{}, common.Wrap(fmt.Errorf("makeCreateUserRequest"), err)
	}

	return request, nil
}

func MakeGetUserRequest(c GinI) (request GetUserRequest, err error) {
	userToGetID, err := getIntFromContext(c, "ID")
	if err != nil {
		return GetUserRequest{}, common.Wrap(fmt.Errorf("makeGetUserRequest"), err)
	}

	request.ID = userToGetID

	if err = request.Validate(); err != nil {
		return GetUserRequest{}, common.Wrap(fmt.Errorf("makeGetUserRequest"), err)
	}

	return request, nil
}

func MakeUpdateUserRequest(c GinI) (request UpdateUserRequest, err error) {
	if err = c.ShouldBindJSON(&request); err != nil {
		return UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), customErrors.ErrBindingRequest)
	}

	userToUpdateID, err := getIntFromContext(c, "ID")
	if err != nil {
		return UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), err)
	}

	request.ID = userToUpdateID

	if err = request.Validate(); err != nil {
		return UpdateUserRequest{}, common.Wrap(fmt.Errorf("makeUpdateUserRequest"), err)
	}

	return request, nil
}

func MakeDeleteUserRequest(c GinI) (request DeleteUserRequest, err error) {
	userToDeleteID, err := getIntFromContext(c, "ID")
	if err != nil {
		return DeleteUserRequest{}, common.Wrap(fmt.Errorf("makeDeleteUserRequest"), err)
	}

	request.ID = userToDeleteID

	if err = request.Validate(); err != nil {
		return DeleteUserRequest{}, common.Wrap(fmt.Errorf("makeDeleteUserRequest"), err)
	}

	return request, nil
}

// - Helpers

func getIntFromContext(c GinI, key string) (int, error) {
	value := c.GetInt(key)
	if value == 0 {
		return 0, fmt.Errorf("error getting %s from context", key)
	}
	return value, nil
}
