package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codec interface {
}

type CodecHandler struct {
}

func (codec *CodecHandler) FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User {
	return models.User{
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
	}
}
