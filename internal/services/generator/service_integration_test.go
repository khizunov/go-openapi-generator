package generator_test

import (
	"os"
	"path/filepath"
	"testing"

	_ "github.com/khizunov/go-openapi-generator/internal/generators/goresty"
	"github.com/khizunov/go-openapi-generator/internal/services/generator"
	"github.com/khizunov/go-openapi-generator/internal/services/parser"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateFromV3_Petstore(t *testing.T) {
	outDir := t.TempDir()
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	p := parser.NewParser(log)
	doc, err := p.ParseV3("../../../testdata/petstore.yaml")
	require.NoError(t, err)

	g := generator.NewGenerator(log)
	err = g.GenerateFromV3(doc, "go-resty", outDir)
	require.NoError(t, err)

	// Model files exist
	assert.FileExists(t, filepath.Join(outDir, "model", "pet.go"))
	assert.FileExists(t, filepath.Join(outDir, "model", "error.go"))

	// API file for pets tag
	assert.FileExists(t, filepath.Join(outDir, "api", "pets.go"))

	// Model file contains the struct
	petContent, err := os.ReadFile(filepath.Join(outDir, "model", "pet.go"))
	require.NoError(t, err)
	assert.Contains(t, string(petContent), "type Pet struct")
	assert.Contains(t, string(petContent), "ID int64")

	// API file contains the method
	apiContent, err := os.ReadFile(filepath.Join(outDir, "api", "pets.go"))
	require.NoError(t, err)
	assert.Contains(t, string(apiContent), "func (c *PetsClient) ListPets")
	assert.Contains(t, string(apiContent), "func (c *PetsClient) ShowPetByID")
}
