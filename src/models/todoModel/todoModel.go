package todoModel

import (
	"fmt"
	"todo-backend-go/src/db"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
	UserId    int
}

func GetTodos() (todos []Todo, err error) {
	err = db.GetORM().Table("todos").Find(&todos).Error
	fmt.Println(todos)
	return
}
