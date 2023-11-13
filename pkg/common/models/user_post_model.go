package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
)

/*---------------------------------------------------------------------------
// Particular Models are a key part of the application, they work as Business
// Objects and contain some of the logic of the app.
//----------------------*/

func (up *UserPost) Create() error {
	userPost, err := up.Repository.CreateUserPost(*up)
	if err != nil {
		return common.Wrap("up.Repository.CreateUserPost", err)
	}
	*up = userPost
	return nil
}
