package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

/*---------------------------------------------------------------------------
// Particular Models are a key part of the application, they work as Business
// Objects and contain some of the logic of the app.
//----------------------*/

func (u *User) GenerateTokenString() (string, error) {
	return auth.GenerateToken(u.ID, u.Username, u.Email, u.GetRole(), u.Config.Auth.JWTSecret, u.Config.Auth.SessionDurationDays)
}

func (u *User) Create() (err error) {
	*u, err = u.Repository.CreateUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.CreateUser", err)
	}
	return nil
}

func (u *User) Get(opts ...options.QueryOption) (err error) {
	*u, err = u.Repository.GetUser(*u, opts...)
	if err != nil {
		return common.Wrap("u.Repository.GetUser", err)
	}
	return nil
}

func (u *User) Update() (err error) {
	*u, err = u.Repository.UpdateUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.UpdateUser", err)
	}
	return nil
}

func (u *User) UpdatePassword() (err error) {
	if err = u.Repository.UpdatePassword(u.ID, u.Password); err != nil {
		return common.Wrap("u.Repository.UpdatePassword", err)
	}
	return nil
}

func (u *User) Delete() (err error) {
	*u, err = u.Repository.DeleteUser(*u)
	if err != nil {
		return common.Wrap("u.Repository.DeleteUser", err)
	}
	return nil
}

func (u *User) Search(page, perPage int, opts ...options.QueryOption) (Users, error) {
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
	u.Password = common.Hash(u.Password, u.Config.Auth.HashSalt)
}

func (u *User) PasswordMatches(password string) bool {
	return u.Password == common.Hash(password, u.Config.Auth.HashSalt)
}

func (u *User) GetRole() auth.Role {
	if u.IsAdmin {
		return auth.AdminRole
	}
	return auth.UserRole
}
