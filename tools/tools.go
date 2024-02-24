//go:build tools
// +build tools

package tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

//go:generate go install github.com/abice/go-enum
//go:generate go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
//go:generate go install mvdan.cc/gofumpt@latest
//go:generate go install github.com/daixiang0/gci@latest

// nolint
import (
	// go enum generator
	_ "github.com/abice/go-enum"
	// gci
	_ "github.com/daixiang0/gci"
	// golangci-lint
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	// gofumpt
	_ "mvdan.cc/gofumpt"
)
