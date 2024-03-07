package common

import "fmt"

type Error struct {
	WrappedError error
	Message      string
	HTTPStatus   int
}

func NewError(msg string, httpStatus int) *Error {
	return &Error{
		WrappedError: fmt.Errorf(msg),
		Message:      msg,
		HTTPStatus:   httpStatus,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	return e.HTTPStatus
}

var (

	// - Transport errors
	ErrUnknown            = NewError("error unknown", 500)
	ErrTooManyRequests    = NewError("error, too many server requests", 429)
	ErrUnauthorized       = NewError("error, unauthorized", 401)
	ErrAllFieldsRequired  = NewError("error, all fields required", 400)
	ErrPasswordsDontMatch = NewError("error, passwords don't match", 400)
	ErrBindingRequest     = NewError("error binding request", 400)
	ErrValidatingRequest  = NewError("error validating request", 400)
	ErrInvalidEmailFormat = NewError("error, invalid email format", 400)
	ErrInvalidValue       = func(field string) error {
		return NewError(fmt.Sprintf("error, invalid value for field %s", field), 400)
	}
	ErrInvalidUsernameLength = func(min, max int) error {
		return NewError(fmt.Sprintf("error, username must contain between %d and %d characters", min, max), 400)
	}
	ErrInvalidPasswordLength = func(min, max int) error {
		return NewError(fmt.Sprintf("error, password must contain between %d and %d characters", min, max), 400)
	}

	// - Service & Repository errors
	ErrInDBTransaction = NewError("error in database transaction", 500)

	// --- Users
	ErrCreatingUser                = NewError("error creating user", 500)
	ErrGettingUser                 = NewError("error getting user", 500)
	ErrUpdatingUser                = NewError("error updating user", 500)
	ErrUpdatingUserDetail          = NewError("error updating user detail", 500)
	ErrDeletingUser                = NewError("error deleting user", 500)
	ErrSearchingUsers              = NewError("error searching users", 500)
	ErrUserNotFound                = NewError("error, user not found", 404)
	ErrUserAlreadyDeleted          = NewError("error, user already deleted", 404)
	ErrUsernameOrEmailAlreadyInUse = NewError("error, username or email already in use", 409)
	ErrWrongPassword               = NewError("error, wrong password", 401)

	// --- User Posts
	ErrCreatingUserPost = NewError("error creating user post", 500)
)
