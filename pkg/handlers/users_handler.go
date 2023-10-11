package handlers

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

type RepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...utils.QueryOption) (models.User, error)
	UserExists(email, username string, opts ...utils.QueryOption) bool
	DeleteUser(id int) (models.User, error)
}

type UserI interface {
	OverwriteFields(username, email string) models.User
	PasswordMatches(password string) bool
	ToEntity() entities.User
	GetAuthRole() entities.Role
}

type UserHandlerI interface {
	Create(r repository.RepositoryLayer) error
}

type UserHandler struct {
	User UserI
}

func New(user UserI) *UserHandler {
	return &UserHandler{User: user}
}

func (h *UserHandler) Create(r repository.RepositoryLayer) error {
	userFromDB, err := r.CreateUser(h.User.(models.User))
	if err != nil {
		return err
	}

	h.User = userFromDB
	return nil
}
