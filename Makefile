# Description: Makefile for go project
projectname ?= go-openapi-generator

.PHONY: lint
lint:
	golangci-lint run ./...

PHONY: test
test: ## run tests with coverage
	mkdir -p build
	go test -v -race $(shell go list ./internal/...) -v -coverprofile=build/coverage.out
	go tool cover -func=build/coverage.out

.PHONY: build
build: ## build the binary
	go build -o bin/$(projectname) ./cmd/app/app.go

.PHONY: go-generate
go-generate:
	go generate ./...

PHONY: fmt
fmt: ## format go files
	gofumpt -w .
	gci write .

.PHONY: tools
tools: ## install build deps
	go generate -tags tools tools/tools.go

.PHONY: clean
clean:
	rm -rf bin/