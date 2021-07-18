package main

import (
	"log"

	"github.com/juank11memphis/go-serverless-boilerplate/internal/core/appcontainer"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/core/database"
	"github.com/juank11memphis/go-serverless-boilerplate/internal/core/routing"
)

func main() {
	log.Println("Locally running....")

	dbOrm, err := database.Connect()
	if err != nil {
		log.Fatal("Unexpected error connecting to local database", err)
		panic(err)
	}

	err = appcontainer.Init(dbOrm)

	if err != nil {
		log.Fatal("Unexpected error initializing app container", err)
		panic(err)
	}

	engine := routing.Build()
	routing.AddGraphqlRouter(engine)
	engine.Run()
}
