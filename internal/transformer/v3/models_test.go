package v3_test

import (
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	transformerv3 "github.com/khizunov/go-openapi-generator/internal/transformer/v3"
	"github.com/pb33f/libopenapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const petStoreModelsYAML = `
openapi: "3.0.0"
info:
  title: Test
  version: "1.0"
paths: {}
components:
  schemas:
    Pet:
      type: object
      required:
        - id
        - name
      properties:
        id:
          type: integer
          format: int64
          description: Unique pet ID
        name:
          type: string
          description: Pet name
        tag:
          type: string
    Error:
      type: object
      properties:
        code:
          type: integer
          format: int32
        message:
          type: string
`

// passthroughGen is a minimal Generator for testing: leaves names unchanged.
type passthroughGen struct{ codegen.Generator }

func (p passthroughGen) ToModelName(n string) string     { return n }
func (p passthroughGen) ToPropertyName(n string) string  { return n }
func (p passthroughGen) ToOperationName(n string) string { return n }
func (p passthroughGen) ToParamName(n string) string     { return n }
func (p passthroughGen) MapType(t, f string) string {
	if t == "integer" && f == "int64" {
		return "int64"
	}
	if t == "integer" {
		return "int"
	}
	return t
}

func TestTransformModels_BasicStruct(t *testing.T) {
	doc, err := libopenapi.NewDocument([]byte(petStoreModelsYAML))
	require.NoError(t, err)
	v3m, errs := doc.BuildV3Model()
	require.Empty(t, errs)

	models, err := transformerv3.TransformModels(&v3m.Model, passthroughGen{})
	require.NoError(t, err)

	names := make([]string, len(models))
	for i, m := range models {
		names[i] = m.Name
	}
	assert.ElementsMatch(t, []string{"Pet", "Error"}, names)

	var pet codegen.Model
	for _, m := range models {
		if m.Name == "Pet" {
			pet = m
		}
	}
	require.Equal(t, "Pet", pet.Name)
	assert.Len(t, pet.Properties, 3)

	propByName := map[string]codegen.Property{}
	for _, p := range pet.Properties {
		propByName[p.OriginalName] = p
	}

	assert.Equal(t, "int64", propByName["id"].Type)
	assert.True(t, propByName["id"].Required)
	assert.Equal(t, "string", propByName["name"].Type)
	assert.True(t, propByName["name"].Required)
	assert.False(t, propByName["tag"].Required)
}
