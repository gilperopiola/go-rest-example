package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
)

/*------------------------------------------------------------------------
// Here we have the business-object part of the models, their behaviour.
//----------------------*/

func (up *UserPost) Create() (err error) {
	*up, err = up.Repository.CreateUserPost(*up)
	if err != nil {
		return common.Wrap("up.Repository.CreateUserPost", err)
	}
	return nil
}
