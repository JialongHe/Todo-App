package routes_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"todo-app/db"
	"todo-app/models"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SetupTestDB sets up a MongoDB test client and overrides the global db.ToDoCollection
func setupTestDB(t *testing.T) *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	testDB := client.Database("todo_test")
	collection := testDB.Collection("todos_test")

	// Override global collection
	db.ToDoCollection = collection

	// Clear data before each test
	err = collection.Drop(context.TODO())
	assert.NoError(t, err)

	return collection
}

func TestCreateToDoRoute(t *testing.T) {
	setupTestDB(t)

	router := gin.Default()
	routes.RegisterToDoRoutes(router)

	// Mock input
	todo := models.ToDo{
		Title:       "Test Todo",
		Description: "Test Description",
		DueDate:     time.Now(),
	}
	body, _ := json.Marshal(todo)

	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var created models.ToDo
	err := json.Unmarshal(resp.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.Equal(t, todo.Title, created.Title)
}

func TestGetToDoByIDRoute(t *testing.T) {
	collection := setupTestDB(t)
	router := gin.Default()
	routes.RegisterToDoRoutes(router)

	// Insert test todo
	todo := models.ToDo{
		Title:       "Sample",
		Description: "Sample Desc",
		DueDate:     time.Now(),
	}
	inserted, err := collection.InsertOne(context.TODO(), todo)
	assert.NoError(t, err)

	id := inserted.InsertedID.(primitive.ObjectID).Hex()
	req, _ := http.NewRequest("GET", "/todos/"+id, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var fetched models.ToDo
	err = json.Unmarshal(resp.Body.Bytes(), &fetched)
	assert.NoError(t, err)
	assert.Equal(t, todo.Title, fetched.Title)
}

func TestUpdateToDoRoute(t *testing.T) {
	collection := setupTestDB(t)
	router := gin.Default()
	routes.RegisterToDoRoutes(router)

	todo := models.ToDo{
		Title:       "Before Update",
		Description: "Before",
		DueDate:     time.Now(),
	}
	result, _ := collection.InsertOne(context.TODO(), todo)
	id := result.InsertedID.(primitive.ObjectID).Hex()

	updated := models.ToDo{
		Title:       "After Update",
		Description: "After",
		DueDate:     time.Now(),
	}
	body, _ := json.Marshal(updated)

	req, _ := http.NewRequest("PUT", "/todos/"+id, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteToDoRoute(t *testing.T) {
	collection := setupTestDB(t)
	router := gin.Default()
	routes.RegisterToDoRoutes(router)

	todo := models.ToDo{
		Title:       "To Be Deleted",
		Description: "Delete me",
		DueDate:     time.Now(),
	}
	result, _ := collection.InsertOne(context.TODO(), todo)
	id := result.InsertedID.(primitive.ObjectID).Hex()

	req, _ := http.NewRequest("DELETE", "/todos/"+id, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	// Make sure it's gone
	var found models.ToDo
	err := collection.FindOne(context.TODO(), todo).Decode(&found)
	assert.Error(t, err) // should be not found
}

func TestGetAllToDosRoute(t *testing.T) {
	collection := setupTestDB(t)
	router := gin.Default()
	routes.RegisterToDoRoutes(router)

	todos := []interface{}{
		models.ToDo{Title: "T1", Description: "D1", DueDate: time.Now()},
		models.ToDo{Title: "T2", Description: "D2", DueDate: time.Now()},
	}
	_, err := collection.InsertMany(context.TODO(), todos)
	assert.NoError(t, err)

	req, _ := http.NewRequest("GET", "/todos?page=1&limit=2", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Results []models.ToDo `json:"results"`
		Count   int           `json:"count"`
	}
	err = json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, response.Count, 2)
	assert.Len(t, response.Results, 2)
}
