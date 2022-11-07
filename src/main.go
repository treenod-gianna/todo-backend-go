package main

import (
	"net/http"
	"todo-backend-go/src/controllers/todoCtrl"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/todos", todoCtrl.Todos)
	r.GET("/todos/:id", todoCtrl.Todo)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
