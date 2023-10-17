package service

import (
	"fmt"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/service/handlers"
)

//------------------------------
//       CREATE USER POST
//------------------------------

func (s *Service) CreateUserPost(createUserPostRequest requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error) {
	handler := handlers.New(
		models.User{},
		createUserPostRequest.ToUserPostModel(),
	)

	if err := handler.CreatePost(s.Repository); err != nil {
		return responses.CreateUserPostResponse{}, common.Wrap(fmt.Errorf("CreateUserPost: user.CreatePost"), err)
	}

	return responses.CreateUserPostResponse{UserPost: handler.ToUserPostResponseModel()}, nil
}
