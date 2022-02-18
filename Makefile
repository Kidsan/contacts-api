build:
	@CGO_ENABLED=0 go build ./cmd/...

test:
	@go test ./... -v

fmt:
	@go fmt ./...

docker:
	@docker build . -t kidsan/contacts-api:latest
	@docker tag kidsan/contacts-api:latest kidsan/contacts-api:$(shell git rev-parse --short --verify main)

generate:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuffer/contactsapi.proto