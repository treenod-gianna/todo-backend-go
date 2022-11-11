package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-backend-go/src/models/todoModel"
	"todo-backend-go/src/service/todoSvc"
)

func Todos(c *gin.Context) {
	todos, err := todoModel.GetTodos()
	if err != nil {
		panic("Error") // TODO Error 함수 생성
	}
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
	return
}

func Todo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := todoSvc.GetTodo(id)
	if err != nil {
		panic("Error")
	}
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}
