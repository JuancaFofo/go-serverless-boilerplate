package coregraphql

import (
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	"go.uber.org/fx"
)

type AllResolvers struct {
	fx.In

	HelloResolver *hello.HelloResolver
	TodoResolver *todo.TodoResolver
}

type Resolvers struct {
	*AllResolvers
}

func NewResolvers(allResolvers AllResolvers) *Resolvers {
	return &Resolvers{
		AllResolvers: &allResolvers,
	}
}

func (r *Resolvers) Hello() string {
	return r.HelloResolver.SayHello()
}

func (r *Resolvers) AddTodo(args *struct{
	Todo *todo.TodoItem
}) *todo.TodoItem {
	return r.TodoResolver.AddTodo(args.Todo)
}
