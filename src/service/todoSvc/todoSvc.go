package todoSvc

import "todo-backend-go/src/models/todoModel"

func GetTodo(id int) string {
	todos := todoModel.GetTodos()
	return todos[id]
}
