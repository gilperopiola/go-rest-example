package service

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type ErrorsMapperProvider interface {
	Map(err error) error
}

type ErrorsMapper struct{}

func (e ErrorsMapper) Map(err error) error {

	// Signup, Login & Users
	if errors.Is(err, repository.ErrCreatingUser) {
		return utils.JoinErrors(entities.ErrCreatingUser, err)
	}

	if errors.Is(err, repository.ErrGettingUser) {
		return utils.JoinErrors(entities.ErrUserNotFound, err)
	}

	if errors.Is(err, repository.ErrUnknown) {
		return utils.JoinErrors(entities.ErrUserNotFound, err)
	}

	// Default to the original error
	return err
}
