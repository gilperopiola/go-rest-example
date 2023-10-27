package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/auth"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-------------------
//       AUTH
//-------------------

func (u *User) ToAuthModel() auth.User {
	return auth.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		IsAdmin:  u.IsAdmin,
	}
}

func (u *User) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(u.ToAuthModel())
}

//----------------
//     USERS
//----------------

func (u *User) Create(r RepositoryLayer) error {
	user, err := r.CreateUser(*u)
	if err != nil {
		return common.Wrap("User.Create", err)
	}
	*u = user
	return nil
}

func (u *User) Get(r RepositoryLayer, opts ...options.QueryOption) error {
	user, err := r.GetUser(*u, opts...)
	if err != nil {
		return common.Wrap("User.Get", err)
	}
	*u = user
	return nil
}

func (u *User) Update(r RepositoryLayer) error {
	user, err := r.UpdateUser(*u)
	if err != nil {
		return common.Wrap("User.Update", err)
	}
	*u = user
	return nil
}

func (u *User) Delete(r RepositoryLayer) error {
	user, err := r.DeleteUser(u.ID)
	if err != nil {
		return common.Wrap("User.Delete", err)
	}
	*u = user
	return nil
}

func (u *User) Search(r RepositoryLayer, page, perPage int) (Users, error) {
	users, err := r.SearchUsers(u.Username, page, perPage, options.WithDetails)
	if err != nil {
		return []User{}, common.Wrap("User.Search", err)
	}
	return users, nil
}

func (u *User) Exists(r RepositoryLayer) bool {
	return r.UserExists(u.Username, u.Email)
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

//-----------------------
//      USER POSTS
//-----------------------

func (up *UserPost) Create(r RepositoryLayer) error {
	userPost, err := r.CreateUserPost(*up)
	if err != nil {
		return common.Wrap("UserPost.CreateUserPost", err)
	}
	*up = userPost
	return nil
}
