package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-backend-go/src/controllers/todoCtrl"
)

type RouteInfo struct {
	Path    string
	Handler func(*gin.Context)
}

func SetupRouteTable(r *gin.Engine) {
	routePostInfos := []RouteInfo{
		{"/todos", todoCtrl.Todos},
		{"/todo", todoCtrl.Todo},
	}
	for _, routeInfo := range routePostInfos {
		//r.POST(routeInfo.Path, routeInfo.Handler)
		fmt.Println("/pokopoko" + routeInfo.Path + ".php")
		r.GET("/pokopoko"+routeInfo.Path+".php", routeInfo.Handler)
	}
}
