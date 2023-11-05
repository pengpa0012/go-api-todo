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
	c.IndentedJSON(http.StatusOK, todos)
}

func deleteTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func toggleTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)
	router.POST("/addTodo", addTodo)
	router.POST("/deleteTodo", deleteTodo)
	router.PATCH("/toggleTodo", toggleTodo)

	router.Run("localhost:8080")
}