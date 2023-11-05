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

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8080")
}