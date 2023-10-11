package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/entities"
	"github.com/gilperopiola/go-rest-example/pkg/utils"
)

// Models are the representation of the entities in the database.
// If you add a new one, remember to add it to the Purge and Migrate functions in pkg/repository/database.go

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	IsAdmin   bool
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserI interface {
	OverwriteFields(username, email string) User
	PasswordMatches(password string) bool
	ToEntity() entities.User
	GetAuthRole() entities.Role
}

func (user User) OverwriteFields(username, email string) User {
	if username != "" {
		user.Username = username
	}
	if email != "" {
		user.Email = email
	}
	return user
}

func (user User) PasswordMatches(password string) bool {
	return user.Password == utils.Hash(user.Email, password)
}

func (user User) ToEntity() entities.User {
	return entities.User{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		IsAdmin:   user.IsAdmin,
		Deleted:   user.Deleted,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (user User) GetAuthRole() entities.Role {
	if user.IsAdmin {
		return entities.AdminRole
	}
	return entities.UserRole
}
