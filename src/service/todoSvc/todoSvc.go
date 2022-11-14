package todoSvc

import (
	"strconv"
	"todo-backend-go/src/models/todoModel"
)

func AddTodo(title string, userId int) error {
	todo := todoModel.Todo{Title: title, UserId: userId}
	err := todoModel.AddTodo(todo)
	if err != nil {
		return err
	}
	return err
}
func UpdateTodo(todo todoModel.Todo) error {
	//func UpdateTodo(id, userId int, completed bool) error {
	//todo := todoModel.Todo{Id: id, Completed: completed, UserId: userId}
	err := todoModel.UpdateTodo(todo)
	return err
}

func UpdateTodoCsv(todo []string) error {
	//func UpdateTodo(id, userId int, completed bool) error {
	//todo := todoModel.Todo{Id: id, Completed: completed, UserId: userId}
	num, _ := strconv.Atoi(todo[0])
	valBool, _ := strconv.ParseBool(todo[2])
	numId, _ := strconv.Atoi(todo[3])
	var todos = todoModel.Todo{num, todo[1], valBool, numId}
	err := todoModel.UpdateTodo(todos)
	return err
}
