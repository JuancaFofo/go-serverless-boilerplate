package coregraphql

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

const (
	GET  = "get"
	POST = "post"
)

func BuildGinEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors())
	return engine
}

func SetupServer(ginEngine *gin.Engine, resolvers *Resolvers) {
	addGraphqlRouter(ginEngine, resolvers)
}

func addGraphqlRouter(ginEngine *gin.Engine, resolvers *Resolvers) {
	group := ginEngine.Group("/")
	graphqlSchema := CreateSchema(resolvers)
	group.POST("graphql", gin.WrapH(&relay.Handler{
		Schema: graphqlSchema,
	}))
}

func CreateSchema(resolvers *Resolvers) *graphql.Schema {
	return graphql.MustParseSchema(AppSchema, resolvers)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		// Handle browser preflight requests, where it asks for allowed origin.
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
