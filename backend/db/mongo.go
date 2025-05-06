package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ToDoCollection *mongo.Collection

func InitMongo() {
	// Get Mongo URI and DB name from environment
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("MONGO_DB_NAME")

	// Create a new MongoDB client and connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	// TodoCollection is the collection where we will store our todos
	ToDoCollection = client.Database(dbName).Collection("todos")
	log.Println("Connected to MongoDB.")
}
