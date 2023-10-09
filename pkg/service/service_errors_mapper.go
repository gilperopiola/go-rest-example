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
		return utils.Wrap(err, entities.ErrCreatingUser)
	}

	if errors.Is(err, repository.ErrUserNotFound) {
		return utils.Wrap(err, entities.ErrUserNotFound)
	}

	if errors.Is(err, repository.ErrUserAlreadyDeleted) {
		return utils.Wrap(err, entities.ErrUserNotFound)
	}

	// General errors
	if errors.Is(err, repository.ErrUnknown) {
		return utils.Wrap(err, entities.ErrUnknown)
	}

	// Default to ErrUnknown
	return entities.ErrUnknown
}
