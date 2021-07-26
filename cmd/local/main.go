package main

import (
	"log"

	corecontainer "github.com/juank11memphis/go-serverless-boilerplate/internal/core/container"
	coredatabase "github.com/juank11memphis/go-serverless-boilerplate/internal/core/database"
	coregraphql "github.com/juank11memphis/go-serverless-boilerplate/internal/core/graphql"
)

func main() {
	log.Println("Locally running....")

	_, err := coredatabase.Connect()
	if err != nil {
		log.Fatal("Unexpected error connecting to local database", err)
		panic(err)
	}
	ginEngine := coregraphql.BuildGinEngine()
	corecontainer.StartAppContainer(ginEngine)
	ginEngine.Run()
}
