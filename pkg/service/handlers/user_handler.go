package handlers

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

// The UserHandler holds a models.User inside, and handles all the operations related to Users

type UserHandler struct {
	User models.User
	Post models.UserPost
}

func New(user models.User, post models.UserPost) *UserHandler {
	return &UserHandler{User: user, Post: post}
}

func (h *UserHandler) ToResponseModel() responses.User {
	return h.User.ToResponseModel()
}

func (h *UserHandler) ToUserPostResponseModel() responses.UserPost {
	return h.Post.ToResponseModel()
}

func (h *UserHandler) ToAuthEntity() auth.User {
	return auth.User{
		ID:       h.User.ID,
		Email:    h.User.Email,
		Username: h.User.Username,
		IsAdmin:  h.User.IsAdmin,
	}
}

// - Auth

func (h *UserHandler) GetAuthRole() auth.Role {
	if h.User.IsAdmin {
		return auth.AdminRole
	}
	return auth.UserRole
}

func (h *UserHandler) GenerateTokenString(a auth.AuthI) (string, error) {
	return a.GenerateToken(h.ToAuthEntity(), h.GetAuthRole())
}

// - Users

func (h *UserHandler) Create(r repository.RepositoryLayer) (err error) {
	h.User, err = r.CreateUser(h.User)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserHandler.Create"), err)
	}
	return nil
}

func (h *UserHandler) Get(r repository.RepositoryLayer, opts ...repository.QueryOption) (err error) {
	h.User, err = r.GetUser(h.User, opts...)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserHandler.Get"), err)
	}
	return nil
}

func (h *UserHandler) Update(r repository.RepositoryLayer) (err error) {
	h.User, err = r.UpdateUser(h.User)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserHandler.Update"), err)
	}
	return nil
}

func (h *UserHandler) Delete(r repository.RepositoryLayer) (err error) {
	h.User, err = r.DeleteUser(h.User.ID)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserHandler.Delete"), err)
	}
	return nil
}

func (h *UserHandler) Exists(r repository.RepositoryLayer) bool {
	return r.UserExists(h.User.Email, h.User.Username)
}

// - Misc

func (h *UserHandler) HashPassword(salt string) {
	h.User.Password = common.Hash(h.User.Password, salt)
}

func (h *UserHandler) PasswordMatches(password, salt string) bool {
	return h.User.Password == common.Hash(password, salt)
}

func (h *UserHandler) OverwriteFields(username, email, password string) {
	if username != "" {
		h.User.Username = username
	}
	if email != "" {
		h.User.Email = email
	}
	if password != "" {
		h.User.Password = password
	}
}

func (h *UserHandler) OverwriteDetails(firstName, lastName *string) {
	if firstName != nil {
		h.User.Details.FirstName = *firstName
	}
	if lastName != nil {
		h.User.Details.LastName = *lastName
	}
}
