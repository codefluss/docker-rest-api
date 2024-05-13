build:
	@go build -o bin/docker-rest-api

run:
	@go run cmd/main.go -b 0.0.0.0

web-bash:
	@docker compose run --service-ports web bash


test:
	@go test -v ./...