package v3_test

import (
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	transformerv3 "github.com/khizunov/go-openapi-generator/internal/transformer/v3"
	"github.com/pb33f/libopenapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const petStoreOpsYAML = `
openapi: "3.0.0"
info:
  title: Test
  version: "1.0"
paths:
  /pets:
    get:
      operationId: listPets
      summary: List all pets
      tags:
        - pets
      parameters:
        - name: limit
          in: query
          required: false
          schema:
            type: integer
            format: int32
      responses:
        "200":
          description: A list of pets
          content:
            application/json:
              schema:
                type: array
                items:
                  type: object
  /pets/{petId}:
    get:
      operationId: showPetById
      summary: Info for a specific pet
      tags:
        - pets
      parameters:
        - name: petId
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Expected response to a valid request
`

func TestTransformOperations_Basic(t *testing.T) {
	doc, err := libopenapi.NewDocument([]byte(petStoreOpsYAML))
	require.NoError(t, err)
	v3m, errs := doc.BuildV3Model()
	require.Empty(t, errs)

	ops, err := transformerv3.TransformOperations(&v3m.Model, passthroughGen{})
	require.NoError(t, err)
	require.Len(t, ops, 2)

	opByID := map[string]codegen.Operation{}
	for _, op := range ops {
		opByID[op.OriginalID] = op
	}

	list := opByID["listPets"]
	assert.Equal(t, "GET", list.Method)
	assert.Equal(t, "/pets", list.Path)
	assert.Equal(t, []string{"pets"}, list.Tags)
	require.Len(t, list.Parameters, 1)
	assert.Equal(t, "limit", list.Parameters[0].OriginalName)
	assert.Equal(t, codegen.ParameterInQuery, list.Parameters[0].In)
	assert.False(t, list.Parameters[0].Required)

	show := opByID["showPetById"]
	assert.Equal(t, "/pets/{petId}", show.Path)
	require.Len(t, show.Parameters, 1)
	assert.Equal(t, codegen.ParameterInPath, show.Parameters[0].In)
	assert.True(t, show.Parameters[0].Required)
}
