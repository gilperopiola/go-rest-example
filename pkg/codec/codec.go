package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codec struct{}

type CodecIFace interface {

	// From Requests to Entities or Models
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromLoginRequestToUserCredentials(request entities.LoginRequest) entities.UserCredentials

	// From Entities to Models
	FromUserCredentialsToUserModel(userCredentials entities.UserCredentials) models.User

	// From Models to Entities
	FromUserModelToUserCredentialsEntities(model models.User) entities.UserCredentials
	FromUserModelToEntities(model models.User) entities.User
}

func NewCodec() Codec {
	return Codec{}
}
