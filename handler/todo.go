package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafael-almeida/go-todo/model"
	"github.com/rafael-almeida/go-todo/repository"
)

func GetTodos(ctx *gin.Context) {
	todos := repository.Scan()
	ctx.JSON(http.StatusOK, todos)
}

func GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	todo := repository.Get(id)

	if todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "todo " + id + " was not found"})
		return
	}

	ctx.JSON(http.StatusOK, *todo)
}

func PostTodo(ctx *gin.Context) {
	todo := model.Todo{}
	err := ctx.ShouldBind(&todo)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "body does not match Todo"})
		return
	}

	successful := repository.Put(&todo)

	if !successful {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "unable to insert todo"})
		return
	}

	ctx.Status(http.StatusOK)
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo := repository.Delete(id)

	if todo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "todo " + id + " was not found"})
		return
	}

	ctx.JSON(http.StatusOK, *todo)
}
