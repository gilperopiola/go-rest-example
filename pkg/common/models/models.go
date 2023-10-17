package models

import (
	"time"
)

// Models are the representation of the database schema. They are used in the Service & Repository Layers.
// If you add a new one, you should add it to the methods in pkg/repository/database.go

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
}

type UserDetail struct {
	ID        int    `gorm:"primaryKey"`
	UserID    int    `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPost struct {
	ID     int    `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Body   string `gorm:"type:text"`
	UserID int    `gorm:"not null"`
}

type UserPosts []UserPost
