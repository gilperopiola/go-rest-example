package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
)

/*------------------------------------------------------------------------
// Here we have the business-object part of the models, their behaviour.
//----------------------*/

func (u *User) GenerateTokenString() (string, error) {
	return auth.GenerateToken(u.ID, u.Username, u.Email, u.GetRole(), u.Config.JWTSecret, u.Config.SessionDurationDays)
}

// Create creates a new user record in the database
func (u *User) Create() error {
	createdUser, err := u.Repository.CreateUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.CreateUser", err)
	}
	*u = createdUser
	return nil
}

// Get retrieves a user from the database based on the current user's fields
func (u *User) Get(opts ...any) error {
	updatedUser, err := u.Repository.GetUser(*u, opts...)
	if err != nil {
		return common.Wrap("u.Repository.GetUser", err)
	}
	*u = updatedUser
	return nil
}

// Update updates the user's information in the database
func (u *User) Update() error {
	updatedUser, err := u.Repository.UpdateUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.UpdateUser", err)
	}
	*u = updatedUser
	return nil
}

// UpdatePassword updates the user's password.
func (u *User) UpdatePassword() error {
	if err := u.Repository.UpdatePassword(u.ID, u.Password); err != nil {
		return common.Wrap("u.Repository.UpdatePassword", err)
	}
	return nil
}

// Delete marks the user as deleted in the database
func (u *User) Delete() error {
	deletedUser, err := u.Repository.DeleteUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.DeleteUser", err)
	}
	*u = deletedUser
	return nil
}

func (u *User) Search(page, perPage int, opts ...any) (Users, error) {
	users, err := u.Repository.SearchUsers(page, perPage, opts...)
	if err != nil {
		return []User{}, common.Wrap("u.Repository.SearchUsers", err)
	}
	return users, nil
}

func (u *User) OverwriteFields(username, email string) {
	if username != "" {
		u.Username = username
	}
	if email != "" {
		u.Email = email
	}
}

func (u *User) OverwriteDetails(firstName, lastName *string) {
	if firstName != nil {
		u.Details.FirstName = *firstName
	}
	if lastName != nil {
		u.Details.LastName = *lastName
	}
}

func (u *User) HashPassword() {
	u.Password = common.Hash(u.Password, u.Config.HashSalt)
}

func (u *User) PasswordMatches(password string) bool {
	return u.Password == common.Hash(password, u.Config.HashSalt)
}

func (u *User) GetRole() auth.Role {
	if u.IsAdmin {
		return auth.AdminRole
	}
	return auth.UserRole
}
