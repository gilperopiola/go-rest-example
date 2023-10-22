package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
)

func (up *UserPost) Create(r RepositoryLayer) error {
	userPost, err := r.CreateUserPost(*up)
	if err != nil {
		return common.Wrap("UserPost.CreateUserPost", err)
	}
	*up = userPost
	return nil
}
