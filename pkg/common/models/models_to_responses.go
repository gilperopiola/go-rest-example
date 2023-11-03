package models

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
)

/*---------------------------------------------------------------------------
// When the Service layer calls the Repository layer, the output is a Model.
// Here we transform those Models into Response Models, returned on our Custom Responses
// to the Transport layer.
------------------------*/

/*-------------------
//      USERS
//-----------------*/

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

func (u Users) ToResponseModel() []responses.User {
	users := []responses.User{}
	for _, user := range u {
		users = append(users, user.ToResponseModel())
	}
	return users
}

func (u UserDetail) ToResponseModel() responses.UserDetail {
	return responses.UserDetail{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

/*-------------------
//      POSTS
//-----------------*/

func (p UserPost) ToResponseModel() responses.UserPost {
	return responses.UserPost{
		ID:    p.ID,
		Title: p.Title,
		Body:  p.Body,
	}
}

func (p UserPosts) ToResponseModel() []responses.UserPost {
	posts := []responses.UserPost{}
	for _, post := range p {
		posts = append(posts, post.ToResponseModel())
	}
	return posts
}
