package models

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/naaltunian/go-mongo/utils"
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
	ctx := context.TODO()
	var user UserModel

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	user.ID = primitive.NewObjectID()

	db := utils.DB.MongoDb

	collection := db.Collection("users")

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}
