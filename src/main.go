package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json: "id`
	Item      string `json: "title`
	Completed bool   `json: "completed`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "1", Item: "Read Book", Completed: false},
	{ID: "1", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, todos)
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo id not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func main() {
	router := gin.Default()
	router.GET("/todo", getTodos)
	router.GET("/todo/:id", getTodo)
	router.POST("/todo", addTodo)
	router.Run(":6969")
}
