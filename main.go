package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin := gin.Default()

	gin.GET("/hello", Get)

	gin.Run()
}

func Get(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}
