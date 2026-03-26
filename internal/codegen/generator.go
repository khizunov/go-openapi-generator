package codegen

import "io/fs"

// GeneratorType classifies what a generator produces.
type GeneratorType string

const (
	GeneratorTypeClient GeneratorType = "client"
	GeneratorTypeServer GeneratorType = "server"
)

// Option is a user-configurable flag exposed by a generator.
type Option struct {
	Name        string
	Description string
	Default     string
}

// OutputFile pairs a template with a destination path and data.
type OutputFile struct {
	TemplateName string
	OutputPath   string
	Data         any
}

// Generator is the plugin interface every code generator must implement.
type Generator interface {
	Name() string
	Tag() GeneratorType
	Options() []Option

	ToModelName(name string) string
	ToOperationName(name string) string
	ToParamName(name string) string
	ToPropertyName(name string) string

	MapType(openAPIType, format string) string

	ReservedWords() map[string]string

	TemplateFS() fs.FS

	OutputFiles(models []Model, operations []Operation) []OutputFile
}
