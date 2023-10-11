package errors

import "fmt"

type Error struct {
	Message string
	Err     error // This is the wrapped error
}

func New(err error) *Error {
	return &Error{
		Message: err.Error(),
		Err:     err,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Unwrap() error {
	return e.Err
}

var (
	// - General errors
	ErrUnknown           = New(fmt.Errorf("error unknown"))
	ErrNilError          = New(fmt.Errorf("error, unexpected nil error"))
	ErrUnauthorized      = New(fmt.Errorf("error, unauthorized"))
	ErrBindingRequest    = New(fmt.Errorf("error binding request"))
	ErrAllFieldsRequired = New(fmt.Errorf("error, all fields required"))

	// - User errors
	ErrCreatingUser                = New(fmt.Errorf("error creating user"))
	ErrGettingUser                 = New(fmt.Errorf("error getting user"))
	ErrUpdatingUser                = New(fmt.Errorf("error updating user"))
	ErrDeletingUser                = New(fmt.Errorf("error deleting user"))
	ErrUserNotFound                = New(fmt.Errorf("error, user not found"))
	ErrUserAlreadyDeleted          = New(fmt.Errorf("error, user already deleted"))
	ErrUsernameOrEmailAlreadyInUse = New(fmt.Errorf("error, username or email already in use"))

	ErrInvalidEmailFormat    = New(fmt.Errorf("error, invalid email format"))
	ErrInvalidUsernameLength = New(fmt.Errorf("error, username either too short or too long"))
	ErrInvalidPasswordLength = New(fmt.Errorf("error, password either too short or too long"))
	ErrPasswordsDontMatch    = New(fmt.Errorf("error, passwords don't match"))
	ErrWrongPassword         = New(fmt.Errorf("error, wrong password"))
)