generate:
	@go run cmd/gen/main.go

build: generate
	@go build
