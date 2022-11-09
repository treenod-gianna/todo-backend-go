package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo-backend-go/src/models/todoModel"
	"todo-backend-go/src/service/todoSvc"
)

func Todos(c *gin.Context) {
	todos := todoModel.GetTodos()
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func Todo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo := todoSvc.GetTodo(id)
	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}
