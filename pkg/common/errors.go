package common

import "fmt"

type Error struct {
	message string
	err     error // This is the wrapped error
}

func NewError(err error) *Error {
	return &Error{
		message: err.Error(),
		err:     err,
	}
}

func (e *Error) Error() string {
	return e.message
}

var (
	// - General errors
	ErrUnknown           = NewError(fmt.Errorf("error unknown"))
	ErrUnauthorized      = NewError(fmt.Errorf("error, unauthorized"))
	ErrBindingRequest    = NewError(fmt.Errorf("error binding request"))
	ErrAllFieldsRequired = NewError(fmt.Errorf("error, all fields required"))
	ErrInvalidValue      = NewError(fmt.Errorf("error, invalid value"))
	ErrTooManyRequests   = NewError(fmt.Errorf("error, too many server requests"))

	// - User errors
	ErrCreatingUser                = NewError(fmt.Errorf("error creating user"))
	ErrGettingUser                 = NewError(fmt.Errorf("error getting user"))
	ErrUpdatingUser                = NewError(fmt.Errorf("error updating user"))
	ErrDeletingUser                = NewError(fmt.Errorf("error deleting user"))
	ErrSearchingUsers              = NewError(fmt.Errorf("error searching users"))
	ErrUserNotFound                = NewError(fmt.Errorf("error, user not found"))
	ErrUserAlreadyDeleted          = NewError(fmt.Errorf("error, user already deleted"))
	ErrUsernameOrEmailAlreadyInUse = NewError(fmt.Errorf("error, username or email already in use"))
	ErrInvalidEmailFormat          = NewError(fmt.Errorf("error, invalid email format"))
	ErrInvalidUsernameLength       = NewError(fmt.Errorf("error, username either too short or too long"))
	ErrInvalidPasswordLength       = NewError(fmt.Errorf("error, password either too short or too long"))
	ErrPasswordsDontMatch          = NewError(fmt.Errorf("error, passwords don't match"))
	ErrWrongPassword               = NewError(fmt.Errorf("error, wrong password"))

	// - User posts errors
	ErrCreatingUserPost = NewError(fmt.Errorf("error creating user post"))
)
