package service

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type errorsMapperInterface interface {
	Map(err error) error
}

type ErrorsMapper struct{}

func NewErrorsMapper() ErrorsMapper {
	return ErrorsMapper{}
}

func (e ErrorsMapper) Map(err error) error {

	// If we're here we shouldn't have a nil error
	if err == nil {
		return entities.ErrNilError
	}

	// Auth & Users errors
	if errors.Is(err, repository.ErrCreatingUser) {
		return utils.JoinErrors(err, entities.ErrCreatingUser)
	}

	if errors.Is(err, repository.ErrGettingUser) {
		return utils.JoinErrors(err, entities.ErrUserNotFound)
	}

	if errors.Is(err, repository.ErrUserAlreadyDeleted) {
		return utils.JoinErrors(err, entities.ErrUserNotFound)
	}

	// General errors

	if errors.Is(err, repository.ErrUnknown) {
		return utils.JoinErrors(err, entities.ErrUserNotFound)
	}

	// Default to the original error
	return err
}
