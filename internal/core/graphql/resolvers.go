package coregraphql

import (
	"github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"
	"go.uber.org/fx"
)

type AllResolvers struct {
	fx.In

	HelloResolver *hello.HelloResolver
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
