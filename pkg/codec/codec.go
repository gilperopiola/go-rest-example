package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codec struct{}

type CodecIFace interface {

	// From Entities to Models
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromUserCredentialsToUserModel(userCredentials entities.UserCredentials) models.User

	// From Models to Entities
	FromUserModelToUserCredentialsEntities(model models.User) entities.UserCredentials
	FromUserModelToEntities(model models.User) entities.User
}
