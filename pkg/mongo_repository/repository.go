package mongo_repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"go.mongodb.org/mongo-driver/mongo"
)

// Compile time check to ensure repository implements the RepositoryLayer interface
var _ MongoRepositoryLayer = (*mongoRepository)(nil)

type MongoRepositoryLayer interface {
	CreateUser(user models.User) (models.User, error)
	GetUser(user models.User, opts ...any) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	UpdatePassword(userID int, password string) error
	DeleteUser(user models.User) (models.User, error)
	SearchUsers(page, perPage int, opts ...any) (models.Users, error)
	CreateUserPost(post models.UserPost) (models.UserPost, error)
}

type mongoRepository struct {
	*mongoDatabase
	collections mongoCollections
}

func New(database *mongoDatabase) *mongoRepository {
	countersCollection := database.db.Database("go-rest-example").Collection("counters")
	usersCollection := database.db.Database("go-rest-example").Collection("users")

	return &mongoRepository{
		mongoDatabase: database,
		collections: mongoCollections{
			counters: countersCollection,
			users:    usersCollection,
		},
	}
}

type mongoCollections struct {
	counters *mongo.Collection
	users    *mongo.Collection
}
