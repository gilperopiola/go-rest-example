package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
)

//------------------------------
//        USER POSTS
//------------------------------

func (s *Service) CreateUserPost(createUserPostRequest requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error) {
	userPost := createUserPostRequest.ToUserPostModel()

	if err := userPost.Create(s.Repository); err != nil {
		return responses.CreateUserPostResponse{}, common.Wrap(fmt.Errorf("CreateUserPost: user.CreatePost"), err)
	}

	return responses.CreateUserPostResponse{UserPost: userPost.ToResponseModel()}, nil
}
