package todo

type TodoService struct {
	dao *TodoDao
}

func NewTodoService(dao *TodoDao) *TodoService {
	return &TodoService{
		dao,
	}
}

func(s *TodoService) AddTodo(todoItem *TodoItem) *TodoItem {
	return s.dao.AddTodo(todoItem)
}
