package mongo_repository

import (
	"github.com/gilperopiola/go-rest-example/pkg/common/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	*Database
	collections mongoCollections
}

type mongoCollections struct {
	counters *mongo.Collection
	users    *mongo.Collection
}

func New(database *Database, mongoCfg config.Mongo) *mongoRepository {
	return &mongoRepository{
		Database: database,
		collections: mongoCollections{
			counters: database.Database(mongoCfg.DBName).Collection("counters"),
			users:    database.Database(mongoCfg.DBName).Collection("users"),
		},
	}
}
