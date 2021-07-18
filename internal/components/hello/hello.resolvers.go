package hello

import (
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/core/appcontainer"
)

type HelloResolver struct{}

func (_ *HelloResolver) SayHello() string {
	container := appcontainer.GetInstance()
	container.DBOrm.Create(&todo.TodoItem{Description: "test1", Completed: false})
	return "Hello Golang Serverless!!!"
}
