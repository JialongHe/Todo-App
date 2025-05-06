package routes

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-app/db"
	"todo-app/models"
)

// RegisterToDoRoutes registers all endpoints for todos
func RegisterToDoRoutes(router *gin.Engine) {
	todo := router.Group("/todos")
	{
		todo.POST("", createToDo)
		todo.GET("", getAllToDos)
		todo.GET("/:id", getToDoByID)
		todo.PUT("/:id", updateToDo)
		todo.DELETE("/:id", deleteToDo)
	}
}

func createToDo(c *gin.Context) {
	var todo models.ToDo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	todo.ID = primitive.NewObjectID()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := db.ToDoCollection.InsertOne(ctx, todo)
	if err != nil {
		log.Println("Bind error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func getAllToDos(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query params
	search := c.DefaultQuery("q", "")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	sortBy := c.DefaultQuery("sortBy", "due_date")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	// Parse page/limit
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)
	skip := (pageInt - 1) * limitInt
	order := 1
	if sortOrder == "desc" {
		order = -1
	}

	// Search filter
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$or": []bson.M{
				{"title": bson.M{"$regex": search, "$options": "i"}},
				{"description": bson.M{"$regex": search, "$options": "i"}},
			},
		}
	}

	// Sort
	sort := bson.D{{Key: sortBy, Value: order}}

	// Count total number of todos (without pagination)
	totalCount, err := db.ToDoCollection.CountDocuments(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB count failed"})
		return
	}

	// Find with options
	opts := options.Find()
	opts.SetSkip(int64(skip))
	opts.SetLimit(int64(limitInt))
	opts.SetSort(sort)

	cursor, err := db.ToDoCollection.Find(ctx, filter, opts)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "DB query failed"})
		return
	}
	defer cursor.Close(ctx)

	var todos []models.ToDo
	if err := cursor.All(ctx, &todos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decoding error"})
		return
	}

	// Return the todos along with total count for pagination
	c.JSON(http.StatusOK, gin.H{
		"page":    pageInt,
		"limit":   limitInt,
		"count":   totalCount, // Total count of matching todos
		"results": todos,
	})
}

func getToDoByID(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.ToDo
	err = db.ToDoCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&todo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ToDo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func updateToDo(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updated models.ToDo
	if err := c.BindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.ToDoCollection.UpdateOne(ctx, bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
		},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo updated"})
}

func deleteToDo(c *gin.Context) {
	idParam := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = db.ToDoCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ToDo deleted"})
}
