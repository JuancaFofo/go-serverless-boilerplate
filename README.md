![Logo](https://cdn-images-1.medium.com/max/1600/1*OezhU9lHTNCk6O6FCUL5fQ.png)

<h1 align="center">Golang Graphql Serverless Architecture Boilerplate</h1>

### 🏠 [Homepage](https://github.com/juank11memphis/go-serverless-boilerplate)

## Structure

```
.
├── bin (output for go binaries)
├── cmd (Application points of entry)
│   └── handlers (lambda handlers)
│       └── handler.go
│   └── local (local development main file)
│       └── main.go
├── internal (Application code)
│   ├── components (app feature components)
│   │   └── hello
│   │      │── hello.dao.go
│   │      │── hello.resolvers.go
│   │      └── hello.service.go
│   └── core (app core components)
│       │── container (DI container)
│       │   └── container.go
│       │── database (database connection)
│       │   └── database.go
│       └── graphql (graphql setup)
│           │── resolvers.go
│           │── schema.go
│           └── server.go
├── .editorconfig
├── .gitignore
├── docker-compose.yml (local postgres database)
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── serverless.yml
```

## How to deploy

1. Make sure your aws credentials are properly set either via an `~/.aws/credentials` file or by setting the corresponding environment variables.
2. Install serverless by running `npm insta1ll -g serverless`
3. Run `make deploy-dev` or `make deploy-prod` depending on what environment you wish to deploy

## How to run locally

### Pre-requisites

1. Install docker
2. Install `gow`: https://github.com/mitranim/gow. We use this to locally run the app in watch mode.

### Starting the app locally

1. Run `make local`. This will start a local postgres db and run the app locally on watch mode

## How to run unit tests

TBD
