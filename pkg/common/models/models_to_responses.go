package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
)

/*---------------------------------------------------------------------------
// When the Service layer calls the Repository layer, the output is a Model.
// Here we transform those Models into Response Models, returned on our Custom Responses
// to the Transport layer.
------------------------*/

func (u Users) ToResponseModel() []responses.User {
	users := make([]responses.User, len(u))
	for i, user := range u {
		users[i] = user.ToResponseModel()
	}
	return users
}

func (p UserPosts) ToResponseModel() []responses.UserPost {
	posts := make([]responses.UserPost, len(p))
	for i, post := range p {
		posts[i] = post.ToResponseModel()
	}
	return posts
}

/*------------------
//      Users
//----------------*/

func (u User) ToResponseModel() responses.User {
	return responses.User{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		IsAdmin:   u.IsAdmin,
		Details:   u.Details.ToResponseModel(),
		Posts:     u.Posts.ToResponseModel(),
		Deleted:   u.Deleted,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (u UserDetail) ToResponseModel() responses.UserDetail {
	return responses.UserDetail{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

/*-------------------
//      Posts
//-----------------*/

func (p UserPost) ToResponseModel() responses.UserPost {
	return responses.UserPost{
		ID:    p.ID,
		Title: p.Title,
		Body:  p.Body,
	}
}
