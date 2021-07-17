package coregraphql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
)

const AppSchema = `
	type Query {
		hello: String!
	}
`

func CreateSchema() *graphql.Schema {
	return graphql.MustParseSchema(AppSchema, &Resolvers{
		HelloResolver: &hello.HelloResolver{},
	})
}
