package handlers

import (
	"errors"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

// mapError maps repository errors to entities errors
func mapRepositoryError(err error) error {

	if err == nil {
		return nil
	}

	// Auth & Users errors
	if errors.Is(err, repository.ErrCreatingUser) {
		return utils.Wrap(err, entities.ErrCreatingUser)
	}

	if errors.Is(err, repository.ErrGettingUser) {
		return utils.Wrap(err, entities.ErrGettingUser)
	}

	if errors.Is(err, repository.ErrUpdatingUser) {
		return utils.Wrap(err, entities.ErrUpdatingUser)
	}

	if errors.Is(err, repository.ErrUserNotFound) {
		return utils.Wrap(err, entities.ErrUserNotFound)
	}

	if errors.Is(err, repository.ErrUserAlreadyDeleted) {
		return utils.Wrap(err, entities.ErrUserAlreadyDeleted)
	}

	// Default to ErrUnknown
	return utils.Wrap(err, entities.ErrUnknown)
}
