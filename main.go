package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const CONN_ADR = "localhost:8080"

func helloWorld(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello, world!")
}

func main() {
	router := gin.Default()
	router.GET("/", helloWorld)
	router.Run(CONN_ADR)
}
