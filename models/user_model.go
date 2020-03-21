package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserModel for User
type UserModel struct {
	ID             primitive.ObjectId `json:"_id,omitempty" bson:"_id, omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
	Password       string             `json:"password,omitempty" bson:"password,omitempty"`
	LinkedIn       string             `json:"linkedIn,omitempty" bson:"linkedIn,omitempty"`
	GithubUsername string             `json:"githubUsername,omitempty" bson:"githubUsername,omitempty"`
	PersonalSite   string             `json:"personalSite,omitempty" bson:"personalSite,omitempty"`
	Bio            string             `json:"bio,omitempty" bson:"bio,omitempty"`
	DateCreated    time.Time          `json:"dateCreated,omitempty" bson:"dateCreated,omitempty"`
}
