package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-backend-go/src/models/todoModel"
	"todo-backend-go/src/service/todoSvc"
)

const RET_OK = 1

func GetTodos(c *gin.Context) {
	userId := 1 // TODO
	todos, err := todoModel.GetTodos(userId)
	if err != nil {
		panic(err) // TODO Error 함수 생성
	}
	c.JSON(http.StatusOK, gin.H{
		"ret":   RET_OK,
		"todos": todos,
	})
	return
}

func AddTodo(c *gin.Context) {
	err := todoSvc.AddTodo("todo1", 1)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})
}

func UpdateTodo(c *gin.Context) {
	err := todoSvc.UpdateTodo(1, 1, true)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": RET_OK,
	})
}
