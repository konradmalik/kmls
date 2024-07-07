.PHONY: build
build:
	@go build -o bin/kmls

.PHONY: test
test:
	@go test ./...

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: lint
lint:
	@golangci-lint run
