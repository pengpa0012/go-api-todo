package main

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	ID     string  `bson:"_id,omitempty"`
	Title  string  `bson:"title"`
	Completed  bool `bson:"completed"`
}

var todos = []Todo{
	{ID: "1", Title: "Wake up", Completed: false},
	{ID: "2", Title: "Eat", Completed: false},
	{ID: "3", Title: "Work", Completed: false},

}

var todosCollection *mongo.Collection

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
					c.IndentedJSON(http.StatusCreated, "Todo Deleted") 
					return
			}
	}
  c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})
}

func toggleCompleted(c *gin.Context) {
	todoID := c.Param("id")
	for index, todo := range todos {
		if todo.ID == todoID {
				todos[index].Completed = !todos[index].Completed
				c.IndentedJSON(http.StatusCreated, "Todo toggled") 
				return
		}
}
c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo item not found"})
}

func main() {
	router := gin.Default()

	// Initialize MongoDB client and collection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
			panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10)
	err = client.Connect(ctx)
	if err != nil {
			panic(err)
	}
	todosCollection = client.Database("go-todo-db").Collection("todos")

	router.GET("/todos", getTodos)
	router.POST("/addTodo", addTodo)
	router.DELETE("/deleteTodo/:id", deleteTodo)
	router.PATCH("/toggleCompleted/:id", toggleCompleted)

	router.Run("localhost:8080")
}