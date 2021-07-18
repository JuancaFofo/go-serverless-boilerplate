.PHONY: build

# ----------------------------
# DEPLOYMENT
# ----------------------------

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

# ----------------------------
# SETTING UP LOCAL DATABASE
# ----------------------------

setup-local-db: docker-compose-down init
	docker-compose up --build postgres

# ----------------------------
# DOCKER
# ----------------------------

docker-clean-up:
	docker stop postgres || true
	docker stop postgres-seed || true
	docker rm -v $(shell docker ps -a -q -f status=exited) 2>&1 || true
	docker rmi $(shell docker images -f "dangling=true" -q) 2>&1 || true
	docker rm postgres || true
	docker rm postgres-seed || true
	docker volume prune -f

docker-compose-down: docker-clean-up
	docker-compose down --remove-orphans

# ----------------------------
# UTILITY
# ----------------------------

format-all:
	gofmt -w .

init:
	$(call create_network)

define create_network
	docker network create goserverless || true
endef
