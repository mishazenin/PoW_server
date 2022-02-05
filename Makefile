include .env
export $(shell sed 's/=.*//' .env)

default: run-server

init:
	@cp .env.sample .env
	@make mocks

run-server:
	@go run -race ./cmd/server/main.go

run-client:
	@go run -race ./cmd/client/main.go

docker-build:
	@docker build -t pow-tcp-client -f build/client.Dockerfile .
	@docker build -t pow-tcp-server -f build/server.Dockerfile .

docker-run:
	@docker-compose up -d

docker-stop:
	@docker-compose stop

lint:
	@golangci-lint run

mocks:
	@echo "Generating mocks..."
	@go generate ./src/...

test:
	@echo "Running tests..."
	@go test ./src/... -cover -count=1
