package mongo_repository

import (
	"context"
	"fmt"
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/config"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*---------------------------------------------------------------
// If you add a new collection, remember to add it here as well
//-----------------------------------------------------*/

var collectionNames = []string{
	"counters",
	"users",
}

var defaultIDCounters = []interface{}{
	bson.D{{Key: "_id", Value: "user_id"}, {Key: "seq", Value: 1}},
	bson.D{{Key: "_id", Value: "user_post_id"}, {Key: "seq", Value: 0}},
}

/*---------------------------
//     Mongo Database
//-------------------------*/

type Database struct {
	*mongo.Client
}

func NewDatabase() *Database {
	var database Database
	var cfg = common.Cfg

	// Create connection. It's deferred closed in main.go.
	var err error
	if database.Client, err = database.connect(cfg.Database.Mongo); err != nil {
		log.Fatalf("error connecting to mongo database: %v", err)
	}

	database.configure(cfg)

	return &database
}

func (database *Database) DB() *mongo.Client {
	if database == nil {
		return nil
	}
	return database.Client
}

/*---------------------------
//  Connect to DB & Ping it
//-------------------------*/

func (database Database) connect(mongoConfig config.Mongo) (*mongo.Client, error) {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoConfig.ConnectionString))
	if err != nil {
		return nil, fmt.Errorf("error connecting: %w", err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("error pinging: %w", err)
	}

	return client, nil
}

/*--------------------------
//    DB Configuration
//------------------------*/

func (database Database) configure(cfg *config.Config) {
	dbConfig := cfg.Database
	mongoConfig := dbConfig.Mongo
	db := database.Database(mongoConfig.DBName)

	// If we destroy all collections we create them again, adding back the default counters
	// and setting up the indices.
	if dbConfig.Destroy {
		destroyAllCollections(db)
		createNewCollections(db)
		insertCounters(db)
		insertIndices(db)
	} else if dbConfig.Clean {
		cleanAllCollections(db)
		insertCounters(db)
	}

	if dbConfig.AdminInsert {
		insertAdmin(db, cfg)
	}
}

func cleanAllCollections(db *mongo.Database) {
	for _, collectionName := range collectionNames {
		db.Collection(collectionName).DeleteMany(context.Background(), bson.M{})
	}
}

func destroyAllCollections(db *mongo.Database) {
	for _, collectionName := range collectionNames {
		db.Collection(collectionName).Drop(context.Background())
	}
}

func createNewCollections(db *mongo.Database) {
	for _, collectionName := range collectionNames {
		db.CreateCollection(context.Background(), collectionName, options.CreateCollection())
	}
}

func insertCounters(db *mongo.Database) {
	countersCollection := db.Collection("counters")
	for _, counter := range defaultIDCounters {
		if _, err := countersCollection.InsertOne(context.Background(), counter); err != nil {
			log.Print(err)
		}
	}
}

func insertIndices(db *mongo.Database) {
	usernameIndex := mongo.IndexModel{Keys: bson.M{"username": 1}, Options: options.Index().SetUnique(true)}
	if _, err := db.Collection("users").Indexes().CreateOne(context.Background(), usernameIndex); err != nil {
		log.Print(err)
	}

	emailIndex := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
	if _, err := db.Collection("users").Indexes().CreateOne(context.Background(), emailIndex); err != nil {
		log.Print(err)
	}
}

func insertAdmin(db *mongo.Database, cfg *config.Config) {
	admin := makeAdminModel("ferra.main@gmail.com", common.Hash(cfg.Database.AdminPassword, cfg.HashSalt))
	if _, err := db.Collection("users").InsertOne(context.Background(), admin); err != nil {
		log.Print(err)
	}
}

func makeAdminModel(email, password string) *models.User {
	return &models.User{
		ID:       1,
		Username: "admin",
		Email:    email,
		Password: password,
		IsAdmin:  true,
		Details: models.UserDetail{
			ID:     1,
			UserID: 1,
		},
	}
}
