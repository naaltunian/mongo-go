package models

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/naaltunian/go-mongo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserModel for User
type UserModel struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	LinkedIn       string             `json:"linkedIn,omitempty" bson:"linkedIn,omitempty"`
	GithubUsername string             `json:"githubUsername,omitempty" bson:"githubUsername,omitempty"`
	PersonalSite   string             `json:"personalSite,omitempty" bson:"personalSite,omitempty"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	DateCreated    time.Time          `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
}

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := context.TODO()
	var user UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	user.ID = primitive.NewObjectID()
	user.DateCreated = time.Now()

	db := utils.DB.MongoDb

	collection := db.Collection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

// GetUser returns a user from the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := context.TODO()
	var user UserModel
	var params = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	db := utils.DB.MongoDb
	collection := db.Collection("users")

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(user)
}

// GetUsers returns all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := context.TODO()

	var users []UserModel

	db := utils.DB.MongoDb
	collection := db.Collection("users")

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user UserModel
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(users)
}
