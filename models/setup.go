package models

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientOptions = options.Client().ApplyURI("mongodb+srv://" + os.Getenv("dbUser") + ":" + os.Getenv("dbPassword") + "@" + os.Getenv("dbCluster"))

// MongoConnection connects with MongoDB
var MongoConnection = connectionToDB()

func connectionToDB() *mongo.Client {
	// Initialize the connection
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Cannot initiliaze connection with MongoDB: " + err.Error())
	}

	log.Println("Succesfully connnected with MongoDB")
	return client
}

// CheckMongoDBConnection validates connection with DB
func CheckMongoDBConnection() bool {
	log.Println("entered to CheckMongoDBConnection...")
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Cannot ping MongoDB connection: " + err.Error())
		return false
	}

	return true
}
