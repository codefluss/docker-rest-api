build:
	@go build -o bin/docker-rest-api

run:
	@go run cmd/main.go -b 0.0.0.0

web-bash:
	@docker compose run --service-ports web bash

rebuild:
	@docker compose up --build -d

test:
	@go test -v ./...