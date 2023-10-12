package handlers

import (
	"github.com/gilperopiola/go-rest-example/pkg/auth"
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

type UserHandlerI interface {
	ToEntity() entities.User

	// DB Methods
	Create(r repository.RepositoryLayer) error
	Get(r repository.RepositoryLayer, opts ...common.QueryOption) error
	Update(r repository.RepositoryLayer) error
	Delete(r repository.RepositoryLayer) error
	Exists(r repository.RepositoryLayer) bool

	// Auth Methods
	GetAuthRole() entities.Role
	GenerateTokenString(a auth.AuthI) (string, error)

	// Helpers
	HashPassword(password string)
	PasswordMatches(password string) bool
	OverwriteFields(username, email, password string)
}

type UserHandler struct {
	User models.User
}

func New(user models.User) *UserHandler {
	return &UserHandler{User: user}
}

func (h *UserHandler) ToEntity() entities.User {
	return entities.User{
		ID:        h.User.ID,
		Email:     h.User.Email,
		Username:  h.User.Username,
		IsAdmin:   h.User.IsAdmin,
		Deleted:   h.User.Deleted,
		CreatedAt: h.User.CreatedAt,
		UpdatedAt: h.User.UpdatedAt,
	}
}

// Helpers

func (h *UserHandler) HashPassword() {
	h.User.Password = common.Hash(h.User.Email, h.User.Password)
}

func (h *UserHandler) PasswordMatches(password string) bool {
	return h.User.Password == common.Hash(h.User.Email, password)
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
