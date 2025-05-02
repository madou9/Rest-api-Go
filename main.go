package main

import (
	"errors"
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

func geTodo(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string)(*todo, error){
	for i, td := range todos{
		if td.ID == id {
			return &todos[i], nil

		}
	}
	return nil, errors.New("todo not found")
}

func toggleTodoStattus(context *gin.Context){
	id := context.Param("id")
	todo, err := getTodoById((id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func main() {
	router := gin.Default() // Create a default Gin router with logging and recovery middleware.
	router.GET("/todos", getTodos) // Define a GET route for /todos, and attach the getTodos handler.
	router.GET("/todos/:id", geTodo)
	router.PATCH("/todos/:id", toggleTodoStattus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
