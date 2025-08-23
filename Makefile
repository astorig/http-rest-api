.PHONY: build
build:
	go build -v -o ./bin/apiserver ./cmd/apiserver

test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

