package goresty_test

import (
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	"github.com/khizunov/go-openapi-generator/internal/generators/goresty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGoRestyGenerator_Name(t *testing.T) {
	g := goresty.New()
	assert.Equal(t, "go-resty", g.Name())
}

func TestGoRestyGenerator_OutputFiles(t *testing.T) {
	g := goresty.New()
	models := []codegen.Model{
		{Name: "Pet"},
		{Name: "Error"},
	}
	ops := []codegen.Operation{
		{Name: "ListPets", Tags: []string{"pets"}},
		{Name: "ShowPetById", Tags: []string{"pets"}},
		{Name: "Healthcheck", Tags: []string{"health"}},
	}

	files := g.OutputFiles(models, ops)

	// 2 model files + 2 tag files
	assert.Len(t, files, 4)

	paths := make([]string, len(files))
	for i, f := range files {
		paths[i] = f.OutputPath
	}
	assert.Contains(t, paths, "model/pet.go")
	assert.Contains(t, paths, "model/error.go")
	assert.Contains(t, paths, "api/pets.go")
	assert.Contains(t, paths, "api/health.go")
}

func TestGoRestyGenerator_TemplateFS(t *testing.T) {
	g := goresty.New()
	fs := g.TemplateFS()
	require.NotNil(t, fs)

	_, err := fs.Open("model.go.tmpl")
	assert.NoError(t, err, "model.go.tmpl must exist in TemplateFS")

	_, err = fs.Open("api.go.tmpl")
	assert.NoError(t, err, "api.go.tmpl must exist in TemplateFS")
}
