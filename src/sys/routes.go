package sys

import (
	"github.com/gin-gonic/gin"
	"todo-backend-go/src/controllers/todoCtrl"
)

type RouteInfo struct {
	Path    string
	Handler func(*gin.Context)
}

func SetupRouteTable(r *gin.Engine) {
	routePostInfos := []RouteInfo{
		{"/getTodos", todoCtrl.GetTodos},
		{"/addTodo", todoCtrl.AddTodo},
		{"/updateTodo", todoCtrl.UpdateTodo},
	}
	for _, routeInfo := range routePostInfos {
		//r.POST(routeInfo.Path, routeInfo.Handler)
		r.POST("/pokopoko"+routeInfo.Path+".php", routeInfo.Handler)
	}
}
