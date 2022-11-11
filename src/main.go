package main

import (
	"todo-backend-go/src/db"
	"todo-backend-go/src/sys"

	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupDatabase()

	r := gin.Default()
	sys.SetupRouteTable(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
