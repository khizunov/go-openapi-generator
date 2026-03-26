package codegen_test

import (
	"io/fs"
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGenerator struct{ name string }

func (s *stubGenerator) Name() string                                         { return s.name }
func (s *stubGenerator) Tag() codegen.GeneratorType                           { return codegen.GeneratorTypeClient }
func (s *stubGenerator) Options() []codegen.Option                            { return nil }
func (s *stubGenerator) ToModelName(n string) string                          { return n }
func (s *stubGenerator) ToOperationName(n string) string                      { return n }
func (s *stubGenerator) ToParamName(n string) string                          { return n }
func (s *stubGenerator) ToPropertyName(n string) string                       { return n }
func (s *stubGenerator) MapType(t, f string) string                           { return t }
func (s *stubGenerator) ReservedWords() map[string]string                     { return nil }
func (s *stubGenerator) TemplateFS() fs.FS                                    { return nil }
func (s *stubGenerator) OutputFiles([]codegen.Model, []codegen.Operation) []codegen.OutputFile {
	return nil
}

func TestRegistry_RegisterAndLookup(t *testing.T) {
	codegen.Register(&stubGenerator{name: "stub"})

	g, err := codegen.Lookup("stub")
	require.NoError(t, err)
	assert.Equal(t, "stub", g.Name())
}

func TestRegistry_LookupUnknown(t *testing.T) {
	_, err := codegen.Lookup("does-not-exist")
	assert.Error(t, err)
}

func TestRegistry_All(t *testing.T) {
	codegen.Register(&stubGenerator{name: "stub-all"})
	all := codegen.All()
	names := make([]string, len(all))
	for i, g := range all {
		names[i] = g.Name()
	}
	assert.Contains(t, names, "stub-all")
}
