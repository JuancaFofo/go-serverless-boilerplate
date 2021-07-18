package coregraphql

import "github.com/juank11memphis/go-serverless-boilerplate/internal/components/hello"

type Resolvers struct {
	HelloResolver *hello.HelloResolver
}

func (r *Resolvers) Hello() string {
	return r.HelloResolver.SayHello()
}
