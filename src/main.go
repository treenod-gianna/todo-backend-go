package main

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"todo-backend-go/src/controllers/todoCtrl"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
	UserId    int
}

func main() {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var todo *Todo
	db.First(&todo, "id = ?", "1")
	fmt.Println(todo)

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
