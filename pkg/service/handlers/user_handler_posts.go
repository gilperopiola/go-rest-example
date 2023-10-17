package handlers

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

func (h *UserHandler) CreatePost(r repository.RepositoryLayer) (err error) {
	h.Post, err = r.CreateUserPost(h.Post)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserHandler.CreatePost"), err)
	}
	return nil
}
