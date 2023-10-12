package handlers

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

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
