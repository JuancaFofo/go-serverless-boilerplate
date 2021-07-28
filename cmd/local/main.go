package main

import (
	"log"

	corecontainer "github.com/juank11memphis/go-serverless-boilerplate/internal/core/container"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
)

func main() {
	log.Println("Locally running....")
	ginEngine := coregraphql.BuildGinEngine()
	corecontainer.StartAppContainer(ginEngine)
	ginEngine.Run()
}
