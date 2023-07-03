package codec

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codecer interface{}

type Codec struct{}

func (codec *Codec) FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User {
	return models.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  hashedPassword,
		Deleted:   false,
		IsAdmin:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (codec *Codec) FromUserModelToEntities(model models.User) entities.User {
	return entities.User{
		ID:        model.ID,
		Email:     model.Email,
		Username:  model.Username,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
