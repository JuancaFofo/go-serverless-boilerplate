package todo

import "gorm.io/gorm"

type TodoItemBase struct {
	Description string
	Completed   bool
}

type TodoItem struct {
	gorm.Model
	Description string
	Completed   bool
}
func(t *TodoItem) Id() int32 {
	return int32(t.ID)
}
