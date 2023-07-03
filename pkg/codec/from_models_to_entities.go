package codec

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (codec *Codec) FromUserModelToEntities(model models.User) entities.User {
	return entities.User{
		ID:        model.ID,
		Email:     model.Email,
		Username:  model.Username,
		CreatedAt: model.CreatedAt.Format(time.RFC3339),
		UpdatedAt: model.UpdatedAt.Format(time.RFC3339),
	}
}
