.PHONY: build

format-all:
	gofmt -w .

build:
	go env -w GO111MODULE=on
	go env -w GOOS=linux
	go build -ldflags="-s -w" -mod=mod -o bin/main cmd/handlers/handler.go

deploy-dev: build
	sls deploy --stage dev --verbose

deploy-prod: build
	sls deploy --stage prod --verbose

destroy-dev:
	sls remove --stage dev --verbose

destroy-prod:
	sls remove --stage prod --verbose
