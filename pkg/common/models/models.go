package models

import (
	"time"

	"github.com/gilperopiola/go-rest-example/pkg/common/config"
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
	CreateUser(user User) (User, error)
	GetUser(user User, opts ...any) (User, error)
	UpdateUser(user User) (User, error)
	UpdatePassword(userID int, password string) error
	DeleteUser(user User) (User, error)
	SearchUsers(page, perPage int, opts ...any) (Users, error)
	CreateUserPost(post UserPost) (UserPost, error)
}

/*-----------------------------------------
// When the Service creates a Model, it passes the Config & Repository to it.
//----------------------------------------------------------------------------*/

type ModelDependencies struct {
	Config     *config.Config `gorm:"-"`
	Repository RepositoryI    `gorm:"-"`
}

/*-----------------------
//       Models
//---------------------*/

type User struct {
	*ModelDependencies `bson:"-"`
	ID                 int        `gorm:"primaryKey" bson:"id"`
	Username           string     `gorm:"unique;not null" bson:"username"`
	Email              string     `gorm:"unique;not null" bson:"email"`
	Password           string     `gorm:"not null" bson:"password"`
	IsAdmin            bool       `bson:"isAdmin"`
	Details            UserDetail `bson:"details"`
	Posts              UserPosts  `gorm:"foreignKey:UserID;references:ID" bson:"posts"`
	Deleted            bool       `bson:"deleted"`
	CreatedAt          time.Time  `bson:"createdAt"`
	UpdatedAt          time.Time  `bson:"updatedAt"`
}

type UserDetail struct {
	ID        int       `gorm:"primaryKey" bson:"id"`
	UserID    int       `gorm:"unique;not null" bson:"-"`
	FirstName string    `gorm:"not null" bson:"firstName"`
	LastName  string    `gorm:"not null" bson:"lastName"`
	CreatedAt time.Time `bson:"-"`
	UpdatedAt time.Time `bson:"-"`
}

type UserPost struct {
	*ModelDependencies `bson:"-"`
	ID                 int    `gorm:"primaryKey" bson:"id"`
	Title              string `gorm:"not null" bson:"title"`
	Body               string `gorm:"type:text" bson:"body"`
	UserID             int    `gorm:"not null" bson:"-"`
}
