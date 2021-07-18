package todo

import "gorm.io/gorm"

type TodoItem struct {
	gorm.Model
	Description string
	Completed   bool
}
