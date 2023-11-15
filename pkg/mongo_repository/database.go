package mongo_repository

import (
	"context"
	"log"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	db *mongo.Client
}

func NewDatabase() *mongoDatabase {
	var database mongoDatabase
	var err error

	// Create connection. It's deferred closed in main.go.
	if database.db, err = database.connectToDB(); err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	database.configure()

	return &database
}

func (database *mongoDatabase) DB() *mongo.Client {
	return database.db
}

/*---------------------------
//  Connect to DB & Ping it
//-------------------------*/

func (database *mongoDatabase) connectToDB() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

/*---------------------------
//    DB Configuration
//-------------------------*/

func (database *mongoDatabase) configure() {
	usersCollection := database.db.Database("go-rest-example").Collection("users")
	countersCollection := database.db.Database("go-rest-example").Collection("counters")

	if err := usersCollection.Drop(context.TODO()); err != nil {
		log.Print(err)
	}
	if err := countersCollection.Drop(context.TODO()); err != nil {
		log.Print(err)
	}

	if false {
		if _, err := usersCollection.DeleteMany(context.TODO(), bson.D{}); err != nil {
			log.Print(err)
		}
		if _, err := countersCollection.DeleteMany(context.TODO(), bson.D{}); err != nil {
			log.Print(err)
		}
	}

	if err := database.db.Database("go-rest-example").CreateCollection(context.TODO(), "users", options.CreateCollection()); err != nil {
		log.Fatal(err)
	}
	if err := database.db.Database("go-rest-example").CreateCollection(context.TODO(), "counters", options.CreateCollection()); err != nil {
		log.Fatal(err)
	}

	usersCollection = database.db.Database("go-rest-example").Collection("users")
	countersCollection = database.db.Database("go-rest-example").Collection("counters")

	// Documents to be inserted
	documents := []interface{}{
		bson.D{{"_id", "user_id"}, {"seq", 1}},
		bson.D{{"_id", "user_post_id"}, {"seq", 0}},
	}

	// Insert documents
	for _, doc := range documents {
		_, err := countersCollection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}
	}

	admin := makeAdminModel("ferra.main@gmail.com", common.Hash(common.Cfg.Database.AdminPassword, common.Cfg.Auth.HashSalt))
	admin.ID = 1
	admin.Details.ID = 1
	if _, err := usersCollection.InsertOne(context.Background(), admin); err != nil {
		log.Fatal(err)
	}

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"username": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}

	indexModel = mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err = usersCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		log.Fatal(err)
	}
}

func makeAdminModel(email, password string) *models.User {
	return &models.User{
		Username: "admin",
		Email:    email,
		Password: password,
		IsAdmin:  true,
	}
}
