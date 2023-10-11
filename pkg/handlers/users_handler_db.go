package handlers

import (
	"github.com/gilperopiola/go-rest-example/pkg/repository"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

func (h *UserHandler) Create(r repository.RepositoryLayer) (err error) {
	h.User, err = r.CreateUser(h.User)
	return mapRepositoryError(err)
}

func (h *UserHandler) Get(r repository.RepositoryLayer, opts ...utils.QueryOption) (err error) {
	h.User, err = r.GetUser(h.User, opts...)
	return mapRepositoryError(err)
}

func (h *UserHandler) Update(r repository.RepositoryLayer) (err error) {
	h.User, err = r.UpdateUser(h.User)
	return mapRepositoryError(err)
}

func (h *UserHandler) Delete(r repository.RepositoryLayer) (err error) {
	h.User, err = r.DeleteUser(h.User.ID)
	return mapRepositoryError(err)
}

func (h *UserHandler) Exists(r repository.RepositoryLayer) bool {
	return r.UserExists(h.User.Email, h.User.Username)
}
