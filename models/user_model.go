package models

import "time"

type UserModel struct {
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	LinkedIn       string    `json:"linkedIn,omitempty"`
	GithubUsername string    `json:"githubUsername,omitempty"`
	PersonalSite   string    `json:"personalSite,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	DateCreated    time.Time `json:"dateCreated,omitempty"`
}
