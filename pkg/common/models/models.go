package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

// Models are the representation of the database schema. They are used in the Service & Repository Layers.
// If you add a new one, you should add it to the methods in pkg/repository/database.go

var AllModels = []interface{}{
	&User{},
	&UserDetail{},
	&UserPost{},
}

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

type RepositoryLayer interface {
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	GetUser(user User, opts ...options.QueryOption) (User, error)
	DeleteUser(id int) (User, error)
	SearchUsers(username string, page, perPage int, opts ...options.PreloadOption) (Users, error)
	UserExists(username, email string, opts ...options.QueryOption) bool

	CreateUserPost(post UserPost) (UserPost, error)
}
