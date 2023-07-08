package service

import (
	"errors"
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type ErrorsMapperIface interface {
	Map(err error) error
}

type ErrorsMapper struct{}

func (e ErrorsMapper) Map(err error) error {

	// Signup
	if errors.Is(err, repository.ErrCreatingUser) {
		return fmt.Errorf("%w:%w", entities.ErrCreatingUser, err)
	}

	// Login
	if errors.Is(err, repository.ErrGettingUser) {
		return fmt.Errorf("%w:%w", entities.ErrUserNotFound, err)
	}

	if errors.Is(err, repository.ErrUnknown) {
		return fmt.Errorf("%w:%w", entities.ErrUserNotFound, err)
	}

	// Default to the original error
	return err
}
