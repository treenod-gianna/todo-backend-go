package todoSvc

import "todo-backend-go/src/models/todoModel"

func AddTodo(title string, userId int) error {
	todo := todoModel.Todo{Title: title, UserId: userId}
	err := todoModel.AddTodo(todo)
	if err != nil {
		return err
	}
	return err
}

func UpdateTodo(id, userId int, completed bool) error {
	todo := todoModel.Todo{Id: id, Completed: completed, UserId: userId}
	err := todoModel.UpdateTodo(todo)
	return err
}
