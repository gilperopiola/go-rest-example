package errors

import "fmt"

type Error struct {
	message string
	err     error // This is the wrapped error
}

func New(err error) *Error {
	return &Error{
		message: err.Error(),
		err:     err,
	}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Unwrap() error {
	return e.err
}

var (
	// - General errors
	ErrUnknown             = New(fmt.Errorf("error unknown"))
	ErrNilError            = New(fmt.Errorf("error, unexpected nil error"))
	ErrUnauthorized        = New(fmt.Errorf("error, unauthorized"))
	ErrBindingRequest      = New(fmt.Errorf("error binding request"))
	ErrAllFieldsRequired   = New(fmt.Errorf("error, all fields required"))
	ErrReadingValueFromCtx = New(fmt.Errorf("error reading value from context"))
	ErrInvalidValue        = New(fmt.Errorf("error, invalid value"))

	// - User errors
	ErrCreatingUser                = New(fmt.Errorf("error creating user"))
	ErrGettingUser                 = New(fmt.Errorf("error getting user"))
	ErrUpdatingUser                = New(fmt.Errorf("error updating user"))
	ErrDeletingUser                = New(fmt.Errorf("error deleting user"))
	ErrSearchingUsers              = New(fmt.Errorf("error searching users"))
	ErrUserNotFound                = New(fmt.Errorf("error, user not found"))
	ErrUserAlreadyDeleted          = New(fmt.Errorf("error, user already deleted"))
	ErrUsernameOrEmailAlreadyInUse = New(fmt.Errorf("error, username or email already in use"))
	ErrInvalidEmailFormat          = New(fmt.Errorf("error, invalid email format"))
	ErrInvalidUsernameLength       = New(fmt.Errorf("error, username either too short or too long"))
	ErrInvalidPasswordLength       = New(fmt.Errorf("error, password either too short or too long"))
	ErrPasswordsDontMatch          = New(fmt.Errorf("error, passwords don't match"))
	ErrWrongPassword               = New(fmt.Errorf("error, wrong password"))

	// - User posts errors
	ErrCreatingUserPost = New(fmt.Errorf("error creating user post"))
)
