package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

/*---------------------------------------------------------------------------
// Particular Models are a key part of the application, they work as business
// objects and contain some of the logic of the app.
//----------------------*/

// We have a RepositoryI here to avoid circular dependencies, our models talk to the repository layer.
type RepositoryI interface {
	CreateUser(user User) (User, error)
	GetUser(user User, opts ...options.QueryOption) (User, error)
	UpdateUser(user User) (User, error)
	UpdatePassword(userID int, password string) error
	DeleteUser(user User) (User, error)
	SearchUsers(page, perPage int, opts ...options.QueryOption) (Users, error)
	UserExists(username, email string, opts ...options.QueryOption) bool

	CreateUserPost(post UserPost) (UserPost, error)
}

/*-------------------
//       AUTH
//-----------------*/

func (u *User) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(u.ID, u.Username, u.Email, u.GetRole())
}

/*----------------
//     USERS
//--------------*/

func (u *User) Create(r RepositoryI) (err error) {
	*u, err = r.CreateUser(*u)
	if err != nil {
		return common.Wrap("r.CreateUser", err)
	}
	return nil
}

func (u *User) Get(r RepositoryI, opts ...options.QueryOption) (err error) {
	*u, err = r.GetUser(*u, opts...)
	if err != nil {
		return common.Wrap("r.GetUser", err)
	}
	return nil
}

func (u *User) Update(r RepositoryI) (err error) {
	*u, err = r.UpdateUser(*u)
	if err != nil {
		return common.Wrap("r.UpdateUser", err)
	}
	return nil
}

func (u *User) UpdatePassword(r RepositoryI) (err error) {
	if err = r.UpdatePassword(u.ID, u.Password); err != nil {
		return common.Wrap("r.UpdatePassword", err)
	}
	return nil
}

func (u *User) Delete(r RepositoryI) (err error) {
	*u, err = r.DeleteUser(*u)
	if err != nil {
		return common.Wrap("r.DeleteUser", err)
	}
	return nil
}

func (u *User) Search(r RepositoryI, page, perPage int, opts ...options.QueryOption) (Users, error) {
	users, err := r.SearchUsers(page, perPage, opts...)
	if err != nil {
		return []User{}, common.Wrap("r.SearchUsers", err)
	}
	return users, nil
}

func (u *User) Exists(r RepositoryI) bool {
	return r.UserExists(u.Username, u.Email)
}

func (u *User) GetRole() auth.Role {
	if u.IsAdmin {
		return auth.AdminRole
	}
	return auth.UserRole
}

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

/*-----------------------
//      USER POSTS
//---------------------*/

func (up *UserPost) Create(r RepositoryI) error {
	userPost, err := r.CreateUserPost(*up)
	if err != nil {
		return common.Wrap("r.CreateUserPost", err)
	}
	*up = userPost
	return nil
}
