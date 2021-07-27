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
			func () *gorm.DB {
				return db
			},
			func () *gin.Engine {
				return ginEngine
			},
			hello.NewHelloResolver,
			hello.NewHelloService,
			hello.NewHelloDao,
			todo.NewTodoDao,
			todo.NewTodoService,
			todo.NewTodoResolver,
			coregraphql.NewResolvers,
		),
		fx.Invoke(coregraphql.SetupServer),
	).Start(context.Background())
}
