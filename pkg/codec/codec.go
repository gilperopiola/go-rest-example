package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type CodecI interface {

	// --- From Requests to Models
	// - Auth
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromLoginRequestToUserModel(request entities.LoginRequest) models.User
	// - Users
	FromCreateUserRequestToUserModel(request entities.CreateUserRequest, hashedPassword string) models.User
	FromGetUserRequestToUserModel(request entities.GetUserRequest) models.User
	FromUpdateUserRequestToUserModel(request entities.UpdateUserRequest) models.User

	// --- From Models to Entities
	FromUserModelToEntities(model models.User) entities.User
}

type Codec struct{}

func NewCodec() *Codec {
	return &Codec{}
}
