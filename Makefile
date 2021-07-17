.PHONY: build

format-all:
	gofmt -w .

build:
	go env -w GO111MODULE=on
	go env -w GOOS=linux
	go build -ldflags="-s -w" -mod=mod -o bin/main cmd/handlers/handler.go

deploy-dev: build
	sls deploy --stage dev --verbose

destroy-dev:
	sls remove --stage dev --verbose
