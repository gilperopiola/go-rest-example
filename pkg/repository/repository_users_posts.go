package repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
)

// CreateUserPost inserts a new post on the database. Title is required, body is optional
func (r *repository) CreateUserPost(post models.UserPost) (models.UserPost, error) {
	if err := r.database.db.Create(&post).Error; err != nil {
		return models.UserPost{}, common.Wrap(err.Error(), common.ErrCreatingUserPost)
	}
	return post, nil
}
