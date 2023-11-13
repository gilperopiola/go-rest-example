package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/repository/options"
)

/*----------------------------------------------------------------------------------------
// Models represent the database schema. They are used in the Service & Repository Layers.
// They are also Business Objects and have methods to interact with the database.
------------------------*/

var AllModels = []interface{}{
	&User{},
	&UserDetail{},
	&UserPost{},
}

type Users []User
type UserPosts []UserPost

/*----------------------------------------------------------------------------------------------------
// We have a RepositoryI here to avoid circular dependencies, our models talk to the repository layer.
------------------------*/

type RepositoryI interface {
	// Users
	CreateUser(user User) (User, error)
	GetUser(user User, opts ...options.QueryOption) (User, error)
	UpdateUser(user User) (User, error)
	UpdatePassword(userID int, password string) error
	DeleteUser(user User) (User, error)
	SearchUsers(page, perPage int, opts ...options.QueryOption) (Users, error)

	// Posts
	CreateUserPost(post UserPost) (UserPost, error)
}

// When the Service creates a Model, it passes the Config & Repository to it.
type ModelDependencies struct {
	Config     *config.Config `gorm:"-"`
	Repository RepositoryI    `gorm:"-"`
}

/*-----------------------
//       MODELS
//---------------------*/

type User struct {
	*ModelDependencies
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
	*ModelDependencies
	ID     int    `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Body   string `gorm:"type:text"`
	UserID int    `gorm:"not null"`
}
