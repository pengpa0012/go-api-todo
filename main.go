package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Completed  bool `json:"completed"`
}

var todos = []todo{
	{ID: "1", Title: "Wake up", Completed: false},
	{ID: "2", Title: "Eat", Completed: false},
	{ID: "3", Title: "Work", Completed: false},

}


func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func addTodo(c *gin.Context) {
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	for _, todo := range todos {
		if todo.ID == newTodo.ID {
				c.JSON(http.StatusConflict, gin.H{"error": "ID already exists"})
				return
		}
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func deleteTodo(c *gin.Context) {
	todoID := c.Param("id")
	for index, todo := range todos {
			if todo.ID == todoID {
					todos = append(todos[:index], todos[index+1:]...)
					c.IndentedJSON(http.StatusNoContent, nil) 
					return
			}
	}
  c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})
}

func toggleTodo(c *gin.Context) {
	// todoID := c.Param("id")
	// find todo and toggle completed
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/addTodo", addTodo)
	router.DELETE("/deleteTodo/:id", deleteTodo)
	router.PATCH("/toggleTodo", toggleTodo)

	router.Run("localhost:8080")
}