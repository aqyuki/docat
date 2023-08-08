fmt:
	@go fmt ./...

test:
	@go test -v ./...

lint:
	@staticcheck ./...

build:
	@go build