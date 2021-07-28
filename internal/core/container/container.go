package corecontainer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/todo"
	coreconfig "github.com/juank11memphis/go-serverless-boilerplate/internal/core/config"
	coredatabase "github.com/juank11memphis/go-serverless-boilerplate/internal/core/database"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
	"go.uber.org/fx"
)

func StartAppContainer(ginEngine *gin.Engine) {
	fx.New(
		fx.Provide(
			// func() *gorm.DB {
			// 	return db
			// },
			func() *gin.Engine {
				return ginEngine
			},
			coreconfig.NewConfig,
			coredatabase.Connect,
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
