package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type Codec struct{}

type CodecInterface interface {

	// From Requests to Models
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromLoginRequestToUserModel(request entities.LoginRequest) models.User
	FromCreateUserRequestToUserModel(request entities.CreateUserRequest, hashedPassword string) models.User
	FromGetUserRequestToUserModel(request entities.GetUserRequest) models.User
	FromUpdateUserRequestToUserModel(request entities.UpdateUserRequest) models.User

	// From Models to Entities
	FromUserModelToEntities(model models.User) entities.User
}

func NewCodec() *Codec {
	return &Codec{}
}

/* ------------------- */
