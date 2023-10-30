package service

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/requests"
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

//-------------------------
//       CREATE USER
//-------------------------

// CreateUser is an admins only endpoint
func (s *service) CreateUser(request *requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	user := request.ToUserModel()

	if user.Exists(s.repository) {
		return responses.CreateUserResponse{}, common.Wrap("CreateUser: user.Exists", common.ErrUsernameOrEmailAlreadyInUse)
	}

	user.HashPassword(s.config.JWT.HashSalt)

	if err := user.Create(s.repository); err != nil {
		return responses.CreateUserResponse{}, common.Wrap("CreateUser: user.Create", err)
	}

	return responses.CreateUserResponse{User: user.ToResponseModel()}, nil
}

//-----------------------
//       GET USER
//-----------------------

func (s *service) GetUser(request *requests.GetUserRequest) (responses.GetUserResponse, error) {
	user := request.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.GetUserResponse{}, common.Wrap("GetUser: user.Get", err)
	}

	return responses.GetUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       UPDATE USER
//--------------------------

func (s *service) UpdateUser(request *requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	user := request.ToUserModel()

	// Check Username/Email availability
	if user.Exists(s.repository) {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Exists", common.ErrUsernameOrEmailAlreadyInUse)
	}

	// Get User
	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Get", err)
	}

	// Overwrite fields that aren't empty
	user.OverwriteFields(request.Username, request.Email, "")
	user.OverwriteDetails(request.FirstName, request.LastName)

	if err := user.Update(s.repository); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Update", err)
	}

	return responses.UpdateUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       DELETE USER
//--------------------------

func (s *service) DeleteUser(request *requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	user := request.ToUserModel()

	// This returns an error if the user is already deleted
	if err := user.Delete(s.repository); err != nil {
		return responses.DeleteUserResponse{}, common.Wrap("DeleteUser: user.Delete", err)
	}

	return responses.DeleteUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//      SEARCH USERS
//--------------------------

// SearchUsers is an admins only endpoint
func (s *service) SearchUsers(request *requests.SearchUsersRequest) (responses.SearchUsersResponse, error) {
	var (
		user    = request.ToUserModel()
		page    = request.Page
		perPage = request.PerPage
	)

	users, err := user.Search(s.repository, page, perPage)
	if err != nil {
		return responses.SearchUsersResponse{}, common.Wrap("SearchUsers: user.Search", err)
	}

	return responses.SearchUsersResponse{
		Users:   users.ToResponseModel(),
		Page:    page,
		PerPage: perPage,
	}, nil
}

//--------------------------
//     CHANGE PASSWORD
//--------------------------

func (s *service) ChangePassword(request *requests.ChangePasswordRequest) (responses.ChangePasswordResponse, error) {
	user := request.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: user.Get", err)
	}

	// Check if old password matches
	if !user.PasswordMatches(request.OldPassword, s.config.JWT.HashSalt) {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: !user.PasswordMatches", common.ErrWrongPassword)
	}

	// Swap passwords
	user.Password = request.NewPassword

	// Hash new password
	user.HashPassword(s.config.JWT.HashSalt)

	if err := user.Update(s.repository); err != nil {
		return responses.ChangePasswordResponse{}, common.Wrap("ChangePassword: user.Update", err)
	}

	return responses.ChangePasswordResponse{
		User: user.ToResponseModel(),
	}, nil
}

//------------------------------
//      CREATE USER POST
//------------------------------

func (s *service) CreateUserPost(request *requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error) {
	userPost := request.ToUserPostModel()

	if err := userPost.Create(s.repository); err != nil {
		return responses.CreateUserPostResponse{}, common.Wrap("CreateUserPost: user.CreatePost", err)
	}

	return responses.CreateUserPostResponse{UserPost: userPost.ToResponseModel()}, nil
}
