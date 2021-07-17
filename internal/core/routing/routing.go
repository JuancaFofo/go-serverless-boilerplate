package routing

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go/relay"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
)

const (
	GET  = "get"
	POST = "post"
)

func Build() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(CORS())
	return engine
}

func AddGraphqlRouter(engine *gin.Engine) {
	group := engine.Group("/")
	group.POST("graphql", gin.WrapH(&relay.Handler{Schema: coregraphql.CreateSchema()}))
}

func CORS() gin.HandlerFunc {
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
