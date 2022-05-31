package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONN_ADR = "localhost:8080"

type todo struct {
	Id string
}

func getTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, todos)
}

func getTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, t := range todos {
		if t.Id == id {
			ctx.JSON(http.StatusOK, t)
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "todo " + id + " was not found"})
}

var todos = []todo{{Id: "1"}, {Id: "2"}}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.Run(CONN_ADR)
}
