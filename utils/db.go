package utils

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type dataBase struct {
// 	session *mongo.Client
// 	users   *mongo.Collection
// }

// Database struct
type Database struct {
	MongoDb *mongo.Database
}

// DB var
var DB Database

// ConnectToDatabase connects to the mongo atlas database
func ConnectToDatabase() {
	ctx := context.TODO()
	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = Database{MongoDb: client.Database("mongo-go")}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected to DB")
}
