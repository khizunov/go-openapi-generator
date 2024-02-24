# Description: Makefile for go project

.PHONY: build
build:
	go build -o bin/ ./cmd/generator/generator.go

.PHONY: go-generate
go-generate:
	go generate ./...