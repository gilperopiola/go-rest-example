package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codec struct{}

type CodecProvider interface {

	// From Requests to Entities or Models
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromLoginRequestToUserModel(request entities.LoginRequest) models.User
	FromGetUserRequestToUserModel(request entities.GetUserRequest) models.User
	FromUpdateUserRequestToUserModel(request entities.UpdateUserRequest) models.User

	// From Models to Entities
	FromUserModelToEntities(model models.User) entities.User
}

func NewCodec() *Codec {
	return &Codec{}
}

/* ------------------- */
