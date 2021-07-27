package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	corecontainer "github.com/juank11memphis/go-serverless-boilerplate/internal/core/container"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
	"gorm.io/gorm"
)

var ginLambda *ginadapter.GinLambda
var db *gorm.DB

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if ginLambda == nil {
		log.Printf("Golang Serverless Backend cold start....")
		ginEngine := coregraphql.BuildGinEngine()
		corecontainer.StartAppContainer(ginEngine, db)
		ginLambda = ginadapter.New(ginEngine)
	}

	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
