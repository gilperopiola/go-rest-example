package mongo_repository

import (
	"context"

	"github.com/gilperopiola/go-rest-example/pkg/common"
	"github.com/gilperopiola/go-rest-example/pkg/common/models"
	mongoOptions "github.com/gilperopiola/go-rest-example/pkg/repository/mongo_repository/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	counterFieldUserID     = "user_id"
	counterFieldUserPostID = "user_post_id"

	optReturnAfter = options.FindOneAndUpdate().SetReturnDocument(options.After)
)

/*-------------------------
//      Create User
//-----------------------*/

func (r *mongoRepository) CreateUser(user models.User) (models.User, error) {
	c := context.Background()

	// Generate new ID, assign it to the user
	newID := r.getNextCounter(counterFieldUserID)
	user.ID = newID
	user.Details.ID = newID

	// Insert the new User. If it fails, subtract the counter
	if _, err := r.collections.users.InsertOne(c, user); err != nil {
		r.subtractCounter(counterFieldUserID)
		return models.User{}, handleMongoError(err)
	}

	return user, nil
}

/*-------------------------
//       Get User
//-----------------------*/

func (r *mongoRepository) GetUser(user models.User, opts ...any) (models.User, error) {
	c := context.Background()

	// This query will find a single User, so we just need 1 of 3 unique identifiers: ID, Username or Email
	user = stripUserIdentifiers(user)
	filter := bson.M{"$or": []bson.M{{"id": user.ID}, {"username": user.Username}, {"email": user.Email}}}

	if err := r.collections.users.FindOne(c, filter).Decode(&user); err != nil {
		return models.User{}, handleMongoError(err)
	}

	return user, nil
}

func stripUserIdentifiers(user models.User) models.User {
	if user.ID != 0 {
		user.Username, user.Email = "", ""
	}
	if user.Username != "" {
		user.Email = ""
	}
	if user.Email != "" {
		user.Username = ""
	}
	return user
}

/*--------------------------------------------
//      Update User -> (replace)
//-----------------------*/

func (r *mongoRepository) UpdateUser(user models.User) (models.User, error) {
	c := context.Background()

	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}

	if _, err := r.collections.users.UpdateOne(c, filter, update); err != nil {
		return models.User{}, handleMongoError(err)
	}

	return user, nil
}

/*-------------------------
//    Update Password
//-----------------------*/

func (r *mongoRepository) UpdatePassword(userID int, newPassword string) error {
	c := context.Background()

	filter := bson.M{"id": userID}
	update := bson.M{"$set": bson.M{"password": newPassword}}

	if _, err := r.collections.users.UpdateOne(c, filter, update); err != nil {
		return handleMongoError(err)
	}

	return nil
}

/*--------------------------------------
//      Delete User -> (soft-delete)
//-----------------------*/

func (r *mongoRepository) DeleteUser(user models.User) (models.User, error) {
	c := context.Background()

	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": bson.M{"deleted": true}}

	if _, err := r.collections.users.UpdateOne(c, filter, update); err != nil {
		return models.User{}, handleMongoError(err)
	}

	return user, nil
}

/*-------------------------
//     Search Users
//-----------------------*/

func (r *mongoRepository) SearchUsers(page, perPage int, opts ...any) (models.Users, error) {
	c := context.Background()

	// Set Filter (if any), Page & PerPage
	filter := mongoOptions.GetFilterFromOptions(opts...)
	findOptions := options.Find().SetSkip(int64(page * perPage)).SetLimit(int64(perPage))

	// Create Cursor to go over the results
	usersCursor, err := r.collections.users.Find(c, filter, findOptions)
	if err != nil {
		return []models.User{}, handleMongoError(err)
	}
	defer usersCursor.Close(c)

	// Iterate over the cursor and decode each document into a models.User
	users, err := getUsersFromCursor(usersCursor)
	if err != nil {
		return []models.User{}, handleMongoError(err)
	}

	return users, nil
}

func getUsersFromCursor(cursor *mongo.Cursor) (models.Users, error) {
	var users models.Users
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return []models.User{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

/*-------------------------
//    Create User Post
//-----------------------*/

func (r *mongoRepository) CreateUserPost(post models.UserPost) (models.UserPost, error) {

	// Get post owner
	user, err := r.GetUser(models.User{ID: post.UserID})
	if err != nil {
		return models.UserPost{}, err
	}

	// Get the next ID, add it to the post, append it to the user's posts
	post.ID = r.getNextCounter(counterFieldUserPostID)
	user.Posts = append(user.Posts, post)

	// Update User
	if _, err := r.UpdateUser(user); err != nil {
		r.subtractCounter(counterFieldUserPostID)
		return models.UserPost{}, err
	}

	return post, nil
}

/*----------------------------------
//  These counters are used to generate AutoIncremental IDs and mimic SQL.
//
//  ID Counters Structure:
//
//    {
// 		  "_id": "user_id",
//  	  "seq": 12
//	  }
//
//    {
// 		  "_id": "user_post_id",
//  	  "seq": 35
//	  }
//	  ...
//
//------------------*/

func handleMongoError(err error) error {
	if err == mongo.ErrNoDocuments {
		return common.Wrap(err.Error(), common.ErrUserNotFound)
	}
	if mongo.IsDuplicateKeyError(err) {
		return common.Wrap(err.Error(), common.ErrUsernameOrEmailAlreadyInUse)
	}
	return common.Wrap(err.Error(), common.ErrUnknown)
}

func (r *mongoRepository) getNextCounter(field string) int {
	c := context.Background()

	filter := bson.M{"_id": field}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	var updatedDoc bson.M
	r.collections.counters.FindOneAndUpdate(c, filter, update, optReturnAfter).Decode(&updatedDoc)

	return int(updatedDoc["seq"].(int32))
}

func (r *mongoRepository) subtractCounter(field string) {
	c := context.Background()
	filter := bson.M{"_id": field}
	update := bson.M{"$inc": bson.M{"seq": -1}}
	r.collections.counters.FindOneAndUpdate(c, filter, update, nil)
}
