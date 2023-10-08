package service

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type errorsMapperI interface {
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
		return utils.WrapErrors(err, entities.ErrCreatingUser)
	}

	if errors.Is(err, repository.ErrGettingUser) {
		return utils.WrapErrors(err, entities.ErrUserNotFound)
	}

	if errors.Is(err, repository.ErrUserAlreadyDeleted) {
		return utils.WrapErrors(err, entities.ErrUserNotFound)
	}

	// General errors
	if errors.Is(err, repository.ErrUnknown) {
		return utils.WrapErrors(err, entities.ErrUserNotFound)
	}

	// Default to the original error
	return err
}
