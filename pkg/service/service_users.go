package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	mongoOptions "github.com/gilperopiola/go-rest-example/pkg/mongo_repository/options"
	"github.com/gilperopiola/go-rest-example/pkg/sql_repository/options"

	"github.com/gin-gonic/gin"
)

/*-----------------------
//       Signup
//---------------------*/

func (s *service) Signup(c *gin.Context, request *requests.SignupRequest) (responses.SignupResponse, error) {
	user := request.ToUserModel(common.Cfg, s.repository)

	if err := user.Create(); err != nil {
		return responses.SignupResponse{}, common.Wrap("Signup: user.Create", err)
	}

	return responses.SignupResponse{User: user.ToResponseModel()}, nil
}

/*---------------------
//       Login
//-------------------*/

func (s *service) Login(c *gin.Context, request *requests.LoginRequest) (responses.LoginResponse, error) {
	user := request.ToUserModel(common.Cfg, s.repository)

	// Get user
	if err := user.Get(options.WithoutDeleted()); err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.Get", err)
	}

	// Check password
	if !user.PasswordMatches(request.Password) {
		return responses.LoginResponse{}, common.Wrap("Login: !user.PasswordMatches", common.ErrWrongPassword)
	}

	// Generate token
	tokenString, err := user.GenerateTokenString()
	if err != nil {
		return responses.LoginResponse{}, common.Wrap("Login: user.GenerateTokenString", common.ErrUnauthorized)
	}

	return responses.LoginResponse{Token: tokenString}, nil
}

/*-------------------------
//      Create User
//-----------------------*/

// CreateUser is an admins only endpoint
func (s *service) CreateUser(c *gin.Context, request *requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	user := request.ToUserModel(common.Cfg, s.repository)

	if err := user.Create(); err != nil {
		return responses.CreateUserResponse{}, common.Wrap("CreateUser: user.Create", err)
	}

	return responses.CreateUserResponse{User: user.ToResponseModel()}, nil
}

/*-----------------------
//       Get User
//---------------------*/

func (s *service) GetUser(c *gin.Context, request *requests.GetUserRequest) (responses.GetUserResponse, error) {
	user := request.ToUserModel(s.repository)

	// Get user (with details & posts)
	if err := user.Get(options.WithDetails(), options.WithPosts()); err != nil {
		return responses.GetUserResponse{}, common.Wrap("GetUser: user.Get", err)
	}

	// If deleted
	if user.Deleted {
		return responses.GetUserResponse{}, common.Wrap("GetUser: user.Deleted", common.ErrUserAlreadyDeleted)
	}

	return responses.GetUserResponse{User: user.ToResponseModel()}, nil
}

/*--------------------------
//      Update User
//------------------------*/

func (s *service) UpdateUser(c *gin.Context, request *requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	user := request.ToUserModel(s.repository)

	// Get user (with details)
	opts := []any{options.WithoutDeleted(), options.WithDetails()}
	if err := user.Get(opts...); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Get", err)
	}

	// Overwrite fields that aren't empty
	user.OverwriteFields(request.Username, request.Email)
	user.OverwriteDetails(request.FirstName, request.LastName)

	// Save
	if err := user.Update(); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Update", err)
	}

	return responses.UpdateUserResponse{User: user.ToResponseModel()}, nil
}

/*--------------------------
//       Delete User
//------------------------*/

func (s *service) DeleteUser(c *gin.Context, request *requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	user := request.ToUserModel(s.repository)

	// Get user
	if err := user.Get(); err != nil {
		return responses.DeleteUserResponse{}, common.Wrap("DeleteUser: user.Get", err)
	}

	// If already deleted
	if user.Deleted {
		return responses.DeleteUserResponse{}, common.Wrap("DeleteUser: user.Deleted", common.ErrUserAlreadyDeleted)
	}

	// If not, soft-delete
	if err := user.Delete(); err != nil {
		return responses.DeleteUserResponse{}, common.Wrap("DeleteUser: user.Delete", err)
	}

	return responses.DeleteUserResponse{User: user.ToResponseModel()}, nil
}

/*--------------------------
//      Search Users
//------------------------*/

// SearchUsers is an admins only endpoint
func (s *service) SearchUsers(c *gin.Context, request *requests.SearchUsersRequest) (responses.SearchUsersResponse, error) {
	var (
		user    = request.ToUserModel(s.repository)
		page    = request.Page
		perPage = request.PerPage
	)

	// Search (with details, filter by username)
	opts := []any{options.WithDetails(), options.WithUsername(user.Username), mongoOptions.WithUsername(user.Username)}
	users, err := user.Search(page, perPage, opts...)
	if err != nil {
		return responses.SearchUsersResponse{}, common.Wrap("SearchUsers: user.Search", err)
	}

	return responses.SearchUsersResponse{Users: users.ToResponseModel(), Page: page, PerPage: perPage}, nil
}

/*-------------------------
//     Change Password
//------------------------*/

func (s *service) ChangePassword(c *gin.Context, request *requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error) {
	user := request.ToUserModel(common.Cfg, s.repository)

	// Get user
	if err := user.Get(options.WithoutDeleted()); err != nil {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: user.Get", err)
	}

	// Check if old password matches
	if !user.PasswordMatches(request.OldPassword) {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: !user.PasswordMatches", common.ErrWrongPassword)
	}

	// Swap passwords, hash new password
	user.Password = request.NewPassword
	user.HashPassword()

	// Save password
	if err := user.UpdatePassword(); err != nil {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: user.UpdatePassword", err)
	}

	return responses.ChangePasswordResponse{User: user.ToResponseModel()}, nil
}

/*------------------------------
//      Create User Post
//----------------------------*/

func (s *service) CreateUserPost(c *gin.Context, request *requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error) {
	userPost := request.ToUserPostModel(s.repository)

	if err := userPost.Create(); err != nil {
		return responses.CreateUserPostResponse{}, common.Wrap("CreateUserPost: userPost.Create", err)
	}

	return responses.CreateUserPostResponse{UserPost: userPost.ToResponseModel()}, nil
}
