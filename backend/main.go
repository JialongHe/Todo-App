package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	"todo-app/db"
	"todo-app/models"
	"todo-app/routes"
)

func prePopulateTodos() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if the collection already has data
	count, err := db.ToDoCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	// If there are no todos, insert default todos
	if count == 0 {
		defaultTodos := []models.ToDo{
			{
				Title:       "Sample Todo 1",
				Description: "This is a sample to-do task.",
				DueDate:     time.Now().Add(24 * time.Hour),
			},
			{
				Title:       "Sample Todo 2",
				Description: "This is another sample task.",
				DueDate:     time.Now().Add(48 * time.Hour),
			},
		}

		// Convert defaultTodos to a slice of interface{}
		todosInterface := make([]interface{}, len(defaultTodos))
		for i, todo := range defaultTodos {
			todosInterface[i] = todo
		}

		_, err := db.ToDoCollection.InsertMany(ctx, todosInterface)
		if err != nil {
			log.Fatalf("Failed to insert default todos: %v", err)
		}

		fmt.Println("Default to-do items have been added to the database.")
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize MongoDB connection
	db.InitMongo()

	// Set up Gin router
	router := gin.Default()
	// Enable CORS middleware
	router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://localhost:3000"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.RegisterToDoRoutes(router)

	// Pre-populate the database with default todos
	prePopulateTodos()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	router.Run(":" + port)
}
