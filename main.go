package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafael-almeida/go-todo/handler"
)

const CONN_ADR = "localhost:8080"

func main() {
	router := gin.Default()

	router.POST("/todos", handler.PostTodo)
	router.GET("/todos", handler.GetTodos)
	router.GET("/todos/:id", handler.GetTodoById)
	router.DELETE("/todos/:id", handler.DeleteTodo)

	router.Run(CONN_ADR)
}
