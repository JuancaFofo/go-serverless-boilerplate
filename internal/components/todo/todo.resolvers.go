package todo

type TodoResolver struct{
	service *TodoService
}

func NewTodoResolver(service *TodoService) *TodoResolver {
	return &TodoResolver{
		service,
	}
}

func(r *TodoResolver) AddTodo(todoItem *TodoItem) *TodoItem {
	return r.service.AddTodo(todoItem)
}
