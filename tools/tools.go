//go:build tools
// +build tools

package tools

// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

//go:generate go install github.com/abice/go-enum

// nolint
import (
	// go enum generator
	_ "github.com/abice/go-enum"
)