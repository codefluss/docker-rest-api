build:
	@go build -o bin/docker-rest-api

run: build
	@./bin/docker-rest-api

web-bash:
	@docker compose run --service-ports web bash

test:
	@go test -v ./...