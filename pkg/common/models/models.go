package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/responses"
)

// Models are the representation of the database schema. They are used in the Service & Repository Layers.
// If you add a new one, remember to add it in pkg/repository/database.go

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
	Details   UserDetail
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ToResponseModel() responses.User {
	return responses.User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		IsAdmin:   u.IsAdmin,
		Details:   u.Details.ToResponseModel(),
		Deleted:   u.Deleted,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type UserDetail struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u UserDetail) ToResponseModel() responses.UserDetail {
	return responses.UserDetail{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
