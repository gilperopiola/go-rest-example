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
func (s *service) CreateUser(createUserRequest requests.CreateUserRequest) (responses.CreateUserResponse, error) {
	user := createUserRequest.ToUserModel()

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

func (s *service) GetUser(getUserRequest requests.GetUserRequest) (responses.GetUserResponse, error) {
	user := getUserRequest.ToUserModel()

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.GetUserResponse{}, common.Wrap("GetUser: user.Get", err)
	}

	return responses.GetUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       UPDATE USER
//--------------------------

func (s *service) UpdateUser(updateUserRequest requests.UpdateUserRequest) (responses.UpdateUserResponse, error) {
	user := updateUserRequest.ToUserModel()

	if user.Exists(s.repository) {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Exists", common.ErrUsernameOrEmailAlreadyInUse)
	}

	if err := user.Get(s.repository, options.WithoutDeleted); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Get", err)
	}

	// Overwrite fields that aren't empty
	user.OverwriteFields(updateUserRequest.Username, updateUserRequest.Email, "")
	user.OverwriteDetails(updateUserRequest.FirstName, updateUserRequest.LastName)

	if err := user.Update(s.repository); err != nil {
		return responses.UpdateUserResponse{}, common.Wrap("UpdateUser: user.Update", err)
	}

	return responses.UpdateUserResponse{User: user.ToResponseModel()}, nil
}

//--------------------------
//       DELETE USER
//--------------------------

func (s *service) DeleteUser(deleteUserRequest requests.DeleteUserRequest) (responses.DeleteUserResponse, error) {
	user := deleteUserRequest.ToUserModel()

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
func (s *service) SearchUsers(searchUsersRequest requests.SearchUsersRequest) (responses.SearchUsersResponse, error) {
	var (
		user    = searchUsersRequest.ToUserModel()
		page    = searchUsersRequest.Page
		perPage = searchUsersRequest.PerPage
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

//------------------------------
//      CREATE USER POST
//------------------------------

func (s *service) CreateUserPost(createUserPostRequest requests.CreateUserPostRequest) (responses.CreateUserPostResponse, error) {
	userPost := createUserPostRequest.ToUserPostModel()

	if err := userPost.Create(s.repository); err != nil {
		return responses.CreateUserPostResponse{}, common.Wrap("CreateUserPost: user.CreatePost", err)
	}

	return responses.CreateUserPostResponse{UserPost: userPost.ToResponseModel()}, nil
}
