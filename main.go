package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct type for each to-do item.
// The struct tags (json:"...") are used for JSON serialization.
type todo struct {
	ID			string `json:"id"`
	Item		string `json:"item"`
	Completed	bool	`json:"completed"`
}

var todos = []todo{
	{ID:"1", Item: "Clean Room", Completed: false},
	{ID:"2", Item: "Read Book", Completed: false},
	{ID:"3", Item: "Record video", Completed: false},
}

// Handler function that responds with all todos as JSON.
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos) // Indented JSON makes the response easier to read in tools like Postman or browsers.
}

// POST a new todo
func addTodo(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default() // Create a default Gin router with logging and recovery middleware.
	router.GET("/todos", getTodos) // Define a GET route for /todos, and attach the getTodos handler.
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
