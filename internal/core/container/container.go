package corecontainer

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
	"go.uber.org/fx"
)

func StartAppContainer(ginEngine *gin.Engine) {
	fx.New(
		// TODO Integrate DB connection with AppContainer
		// fx.Provide(
		// 	func() *DBProxy {
		// 		return &DBProxy{conn: dbConn}
		// 	},
		// ),
		fx.Provide(
			func () *gin.Engine {
				return ginEngine
			},
			hello.NewHelloResolver,
			hello.NewHelloService,
			hello.NewHelloDao,
			coregraphql.NewResolvers,
		),
		fx.Invoke(coregraphql.SetupServer),
	).Start(context.Background())
}
