package codec

import (
	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/models"
)

type CodecIFace interface {
	FromSignupRequestToUserModel(request entities.SignupRequest, hashedPassword string) models.User
	FromUserModelToEntities(model models.User) entities.User
}

type Codec struct{}
