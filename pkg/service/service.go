package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository"

	"github.com/gin-gonic/gin"
)

// Compile time check to validate that the service struct implements the ServiceLayer interface
var _ ServiceLayer = (*service)(nil)

/*----------------------------------------------------------------------------------------
// The ServiceLayer is called by HandleRequest on handlers.go
------------------------*/

type ServiceLayer interface {
	Signup(c *gin.Context, request *requests.SignupRequest) (responses.SignupResponse, error)
	Login(c *gin.Context, request *requests.LoginRequest) (responses.LoginResponse, error)
	CreateUser(c *gin.Context, request *requests.CreateUserRequest) (responses.CreateUserResponse, error)
	GetUser(c *gin.Context, request *requests.GetUserRequest) (responses.GetUserResponse, error)
	UpdateUser(c *gin.Context, request *requests.UpdateUserRequest) (responses.UpdateUserResponse, error)
	DeleteUser(c *gin.Context, request *requests.DeleteUserRequest) (responses.DeleteUserResponse, error)
	SearchUsers(c *gin.Context, request *requests.SearchUsersRequest) (responses.SearchUsersResponse, error)
	ChangePassword(c *gin.Context, request *requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error)
	CreateUserPost(c *gin.Context, request *requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error)
}

type service struct {
	repository repository.RepositoryLayer
}

func New(repository repository.RepositoryLayer) *service {
	return &service{repository: repository}
}
