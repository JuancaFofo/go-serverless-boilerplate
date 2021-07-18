package main

import (
	"log"

	"github.com/juank11memphis/go-serverless-boilerplate/internal/core/routing"
)

func main() {
	log.Printf("Locally running....")
	engine := routing.Build()
	routing.AddGraphqlRouter(engine)
	engine.Run()
}
