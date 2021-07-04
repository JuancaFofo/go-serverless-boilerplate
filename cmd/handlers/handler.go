package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/juank11memphis/jufo-gifts-backend/internal/components/hello"
	"github.com/juank11memphis/jufo-gifts-backend/internal/core/routing"
)

func main() {
	engine := routing.Build()
	handler := routing.CreateLambdaEntrypoint(engine, hello.Path, hello.Method, hello.ProcessRequest)
	lambda.Start(handler)
}
