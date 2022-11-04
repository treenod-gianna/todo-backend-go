package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	todos := []string{"todo1", "todo2", "todo3"}

	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"todos": todos,
		})
	})
	r.GET("/todos/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		todo := todos[id]
		c.JSON(http.StatusOK, gin.H{
			"todo": todo,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
