package models

import (
	"time"
)

/*---------------------------------------------------------------------------
// Models are the representation of the database schema. They are used in the Service & Repository Layers.
// They are probably the most important part of the app.
------------------------*/

var AllModels = []interface{}{
	&User{},
	&UserDetail{},
	&UserPost{},
}

type Users []User

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool
	Details   UserDetail
	Posts     UserPosts `gorm:"foreignKey:UserID;references:ID"`
	Deleted   bool
	CreatedAt time.Time
	UpdatedAt time.Time

	// DTOs
	NewPassword string `gorm:"-"`
}

type UserDetail struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPosts []UserPost

type UserPost struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Body   string `gorm:"type:text"`
	UserID int    `gorm:"not null"`
}
