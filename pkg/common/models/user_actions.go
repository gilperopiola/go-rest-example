package models

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-------------------
//       AUTH
//-------------------

func (u *User) GetAuthRole() auth.Role {
	if u.IsAdmin {
		return auth.AdminRole
	}
	return auth.UserRole
}

func (u *User) ToAuthEntity() auth.User {
	return auth.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		IsAdmin:  u.IsAdmin,
	}
}

func (u *User) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(u.ToAuthEntity(), u.GetAuthRole())
}

//-------------------
//     DATABASE
//-------------------

func (u *User) Create(r RepositoryLayer) error {
	user, err := r.CreateUser(*u)
	if err != nil {
		return common.Wrap(fmt.Errorf("User.Create"), err)
	}
	*u = user
	return nil
}

func (u *User) Get(r RepositoryLayer, opts ...options.QueryOption) error {
	user, err := r.GetUser(*u, opts...)
	if err != nil {
		return common.Wrap(fmt.Errorf("User.Get"), err)
	}
	*u = user
	return nil
}

func (u *User) Update(r RepositoryLayer) error {
	user, err := r.UpdateUser(*u)
	if err != nil {
		return common.Wrap(fmt.Errorf("User.Update"), err)
	}
	*u = user
	return nil
}

func (u *User) Delete(r RepositoryLayer) error {
	user, err := r.DeleteUser(u.ID)
	if err != nil {
		return common.Wrap(fmt.Errorf("User.Delete"), err)
	}
	*u = user
	return nil
}

func (u *User) Search(r RepositoryLayer, page, perPage int) (Users, error) {
	users, err := r.SearchUsers(u.Username, page, perPage, options.WithDetails)
	if err != nil {
		return []User{}, common.Wrap(fmt.Errorf("User.Search"), err)
	}
	return users, nil
}

func (u *User) Exists(r RepositoryLayer) bool {
	return r.UserExists(u.Username, u.Email)
}

//-------------------
//       MISC
//-------------------

func (u *User) HashPassword(salt string) {
	u.Password = common.Hash(u.Password, salt)
}

func (u *User) PasswordMatches(password, salt string) bool {
	return u.Password == common.Hash(password, salt)
}

func (u *User) OverwriteFields(username, email, password string) {
	if username != "" {
		u.Username = username
	}
	if email != "" {
		u.Email = email
	}
	if password != "" {
		u.Password = password
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
