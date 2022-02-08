build:
	@CGO_ENABLED=0 go build .

test:
	@go test ./...

fmt:
	@go fmt ./...

docker:
	@docker build . -t kidsan/contacts-api:latest
	@docker tag kidsan/contacts-api:latest kidsan/contacts-api:$(shell git rev-parse --short --verify main)