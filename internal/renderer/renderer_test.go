package renderer_test

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/khizunov/go-openapi-generator/internal/renderer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender_WritesExpectedFile(t *testing.T) {
	tmplFS := fstest.MapFS{
		"model.go.tmpl": &fstest.MapFile{
			Data: []byte(`package models

// {{ .Name }} is generated.
type {{ .Name }} struct{}
`),
		},
	}

	outDir := t.TempDir()
	outPath := filepath.Join(outDir, "pet.go")

	err := renderer.Render(tmplFS, "model.go.tmpl", map[string]string{"Name": "Pet"}, outPath)
	require.NoError(t, err)

	content, err := os.ReadFile(outPath)
	require.NoError(t, err)
	assert.Contains(t, string(content), "type Pet struct{}")
	assert.Contains(t, string(content), "// Pet is generated.")
}

func TestRender_CreatesDirectories(t *testing.T) {
	tmplFS := fstest.MapFS{
		"api.go.tmpl": &fstest.MapFile{Data: []byte(`package api`)},
	}
	outDir := t.TempDir()
	outPath := filepath.Join(outDir, "nested", "dir", "api.go")

	err := renderer.Render(tmplFS, "api.go.tmpl", nil, outPath)
	require.NoError(t, err)
	assert.FileExists(t, outPath)
}
