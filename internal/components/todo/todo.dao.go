package todo_comp

import "gorm.io/gorm"

type TodoDao struct {
	db *gorm.DB
}

func NewTodoDao(db *gorm.DB) *TodoDao {
	return &TodoDao{
		db,
	}
}

func (d *TodoDao) AddTodo(todoItem *TodoItem) *TodoItem {
	// TODO handle gorm errors
	d.db.Create(todoItem)
	return todoItem
}
