package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

func (codec *Codec) FromUserCredentialsToUserModel(userCredentials entities.UserCredentials) models.User {
	return models.User{
		Email:    userCredentials.Email,
		Username: userCredentials.Username,
		Password: userCredentials.Password,
	}
}
