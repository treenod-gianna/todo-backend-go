package todoSvc

import "todo-backend-go/src/models/todoModel"

func GetTodo(id int) (todoModel.Todo, error) {
	todos, err := todoModel.GetTodos()
	if err != nil {
		return todoModel.Todo{}, err
	}
	return todos[id], err
}
