package common

import "fmt"

type Error struct {
	err     error // This is the wrapped error
	message string
	status  int
}

func NewError(err error, status int) *Error {
	return &Error{
		err:     err,
		message: err.Error(),
		status:  status,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Status() int {
	return e.status
}

var (
	// - Transport errors
	ErrUnknown               = NewError(fmt.Errorf("error unknown"), 500)
	ErrTooManyRequests       = NewError(fmt.Errorf("error, too many server requests"), 429)
	ErrUnauthorized          = NewError(fmt.Errorf("error, unauthorized"), 401)
	ErrAllFieldsRequired     = NewError(fmt.Errorf("error, all fields required"), 400)
	ErrInvalidValue          = NewError(fmt.Errorf("error, invalid value"), 400)
	ErrBindingRequest        = NewError(fmt.Errorf("error binding request"), 400)
	ErrInvalidEmailFormat    = NewError(fmt.Errorf("error, invalid email format"), 400)
	ErrInvalidUsernameLength = NewError(fmt.Errorf("error, username either too short or too long"), 400)
	ErrInvalidPasswordLength = NewError(fmt.Errorf("error, password either too short or too long"), 400)
	ErrPasswordsDontMatch    = NewError(fmt.Errorf("error, passwords don't match"), 400)

	// - Service & Repository errors
	// --- Users
	ErrCreatingUser                = NewError(fmt.Errorf("error creating user"), 500)
	ErrGettingUser                 = NewError(fmt.Errorf("error getting user"), 500)
	ErrUpdatingUser                = NewError(fmt.Errorf("error updating user"), 500)
	ErrUpdatingUserDetail          = NewError(fmt.Errorf("error updating user detail"), 500)
	ErrDeletingUser                = NewError(fmt.Errorf("error deleting user"), 500)
	ErrSearchingUsers              = NewError(fmt.Errorf("error searching users"), 500)
	ErrUserNotFound                = NewError(fmt.Errorf("error, user not found"), 404)
	ErrUserAlreadyDeleted          = NewError(fmt.Errorf("error, user already deleted"), 404)
	ErrUsernameOrEmailAlreadyInUse = NewError(fmt.Errorf("error, username or email already in use"), 409)
	ErrWrongPassword               = NewError(fmt.Errorf("error, wrong password"), 401)

	// --- User Posts
	ErrCreatingUserPost = NewError(fmt.Errorf("error creating user post"), 500)
)
