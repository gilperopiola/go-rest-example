package codec

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

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
