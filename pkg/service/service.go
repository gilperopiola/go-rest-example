package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository"
)

// Compile time check to validate that the service struct implements the ServiceLayer interface
var _ ServiceLayer = (*service)(nil)

/*----------------------------------------------------------------------------------------
// The ServiceLayer is called by HandleRequest on endpoints.go
------------------------*/

type ServiceLayer interface {
	Signup(request *requests.SignupRequest) (responses.SignupResponse, error)
	Login(request *requests.LoginRequest) (responses.LoginResponse, error)
	CreateUser(request *requests.CreateUserRequest) (responses.CreateUserResponse, error)
	GetUser(request *requests.GetUserRequest) (responses.GetUserResponse, error)
	UpdateUser(request *requests.UpdateUserRequest) (responses.UpdateUserResponse, error)
	DeleteUser(request *requests.DeleteUserRequest) (responses.DeleteUserResponse, error)
	SearchUsers(request *requests.SearchUsersRequest) (responses.SearchUsersResponse, error)
	ChangePassword(request *requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error)
	CreateUserPost(request *requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error)
}

type service struct {
	repository repository.RepositoryLayer
}

func New(repository repository.RepositoryLayer) *service {
	return &service{repository: repository}
}
