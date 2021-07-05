package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/juank11memphis/jufo-gifts-backend/internal/components/hello"
	"github.com/juank11memphis/jufo-gifts-backend/internal/core/routing"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		log.Printf("Jufo gifts backend cold start...")
		engine := routing.Build()
		routing.AddRoute(engine, hello.Path, hello.Method, hello.ProcessRequest)
		ginLambda = ginadapter.New(engine)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
