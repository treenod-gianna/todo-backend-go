package todoCtrl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-backend-go/src/service/todoSvc"
)

func Todos(c *gin.Context) {
	todos := todoSvc.GetTodos()
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func Todo(c *gin.Context) {
	//id, _ := strconv.Atoi(c.Param("id"))
	//todo := todos[id]
	//c.JSON(http.StatusOK, gin.H{
	//	"todo": todo,
	//})
}
