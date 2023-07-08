package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (codec *Codec) FromUserModelToUserCredentials(model models.User) entities.UserCredentials {
	return entities.UserCredentials{
		Email:    model.Email,
		Username: model.Username,
		Password: model.Password,
	}
}

func (codec *Codec) FromUserModelToEntities(model models.User) entities.User {
	return entities.User{
		ID:        model.ID,
		Email:     model.Email,
		Username:  model.Username,
		IsAdmin:   model.IsAdmin,
		Deleted:   model.Deleted,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
