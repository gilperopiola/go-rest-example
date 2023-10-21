package models

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
)

func (up *UserPost) Create(r RepositoryLayer) error {
	userPost, err := r.CreateUserPost(*up)
	if err != nil {
		return common.Wrap(fmt.Errorf("UserPost.CreateUserPost"), err)
	}
	*up = userPost
	return nil
}
