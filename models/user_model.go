package models

import (
	"context"
	"encoding/json"
	"io"
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

// UserModels is a slice of UserModel
type userModels []UserModel

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	ctx := context.TODO()
	user := r.Context().Value(UserKey{}).(UserModel)

	user.ID = primitive.NewObjectID()
	user.DateCreated = time.Now()

	db := utils.DB.MongoDb

	collection := db.Collection("users")

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Println(err)
	}

	err = user.toJSON(w)
	if err != nil {
		log.Println("Error marshalling to JSON", err)
	}
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
		log.Println(err)
	}

	err = user.toJSON(w)
}

// GetUsers returns all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ctx := context.TODO()

	var users userModels

	db := utils.DB.MongoDb
	collection := db.Collection("users")

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Println(err)
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var user UserModel
		err := cur.Decode(&user)
		if err != nil {
			log.Println(err)
		}

		users = append(users, user)
	}

	if err := cur.Err(); err != nil {
		log.Println(err)
	}

	err = users.toJSON(w)
}

// FromJSON returns UserModel
func (u *UserModel) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

// toJSON return UserModel as JSON
func (u *UserModel) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// toJSON returns an array of UserModels as JSON
func (u *userModels) toJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

// UserKey for context
type UserKey struct{}

// ValidateUserMiddleware validates UserModel input from the client
func ValidateUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := UserModel{}
		err := user.FromJSON(r.Body)
		if err != nil {
			log.Println("Error deserializing user", err)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey{}, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
