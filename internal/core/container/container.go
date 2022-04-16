package corecontainer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func StartAppContainer(ginEngine *gin.Engine, db *gorm.DB) {
	fx.New(
		fx.Provide(
			func() *gorm.DB {
				return db
			},
			func() *gin.Engine {
				return ginEngine
			},
			hello_comp.NewHelloResolver,
			hello_comp.NewHelloService,
			hello_comp.NewHelloDao,
			todo_comp.NewTodoDao,
			todo_comp.NewTodoService,
			todo_comp.NewTodoResolver,
			coregraphql.NewResolvers,
		),
		fx.Invoke(coregraphql.SetupServer),
	).Start(context.Background())
}
