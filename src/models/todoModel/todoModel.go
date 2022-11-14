package todoModel

import (
	"todo-backend-go/src/db"
)

type Todo struct {
	Id        int    `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	Completed bool   `json:"completed" gorm:"column:completed"`
	UserId    int    `json:"userId" gorm:"column:userId"`
}

func GetTodos(userId int) (todos []Todo, err error) {
	err = db.GetORM().Table("todos").Where("userId = ?", userId).Find(&todos).Error
	return
}

func AddTodo(todo Todo) error {
	err := db.GetORM().Table("todos").Create(&todo).Error
	return err
}

func UpdateTodo(todo Todo) error {
	err := db.GetORM().Table("todos").Where("userId = ? AND id = ?", todo.UserId, todo.Id).
		Update("title", todo.Title).Error
	return err
}
