package goresty

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
)

//go:embed templates
var templateFiles embed.FS

// GoRestyGenerator generates a Go HTTP client using go-resty/resty.
type GoRestyGenerator struct{}

// New returns a new GoRestyGenerator.
func New() *GoRestyGenerator {
	return &GoRestyGenerator{}
}

func (g *GoRestyGenerator) Name() string               { return "go-resty" }
func (g *GoRestyGenerator) Tag() codegen.GeneratorType { return codegen.GeneratorTypeClient }
func (g *GoRestyGenerator) Options() []codegen.Option  { return nil }
func (g *GoRestyGenerator) ReservedWords() map[string]string {
	return map[string]string{
		"break": "break_", "case": "case_", "chan": "chan_",
		"const": "const_", "continue": "continue_", "default": "default_",
		"defer": "defer_", "else": "else_", "fallthrough": "fallthrough_",
		"for": "for_", "func": "func_", "go": "go_",
		"goto": "goto_", "if": "if_", "import": "import_",
		"interface": "interface_", "map": "map_", "package": "package_",
		"range": "range_", "return": "return_", "select": "select_",
		"struct": "struct_", "switch": "switch_", "type": "type_",
		"var": "var_",
	}
}

func (g *GoRestyGenerator) ToModelName(name string) string     { return ToModelName(name) }
func (g *GoRestyGenerator) ToOperationName(name string) string { return ToOperationName(name) }
func (g *GoRestyGenerator) ToParamName(name string) string     { return ToParamName(name) }
func (g *GoRestyGenerator) ToPropertyName(name string) string  { return ToPropertyName(name) }
func (g *GoRestyGenerator) MapType(t, f string) string         { return MapType(t, f) }

func (g *GoRestyGenerator) TemplateFS() fs.FS {
	sub, _ := fs.Sub(templateFiles, "templates")
	return sub
}

type tagOperations struct {
	Tag        string
	Operations []codegen.Operation
}

// OutputFiles returns one output file per model and one per tag.
func (g *GoRestyGenerator) OutputFiles(models []codegen.Model, ops []codegen.Operation) []codegen.OutputFile {
	var files []codegen.OutputFile

	for _, m := range models {
		files = append(files, codegen.OutputFile{
			TemplateName: "model.go.tmpl",
			OutputPath:   "model/" + strings.ToLower(m.Name) + ".go",
			Data:         m,
		})
	}

	tagMap := map[string][]codegen.Operation{}
	tagOrder := []string{}
	for _, op := range ops {
		tag := "default"
		if len(op.Tags) > 0 {
			tag = op.Tags[0]
		}
		if _, exists := tagMap[tag]; !exists {
			tagOrder = append(tagOrder, tag)
		}
		tagMap[tag] = append(tagMap[tag], op)
	}
	for _, tag := range tagOrder {
		files = append(files, codegen.OutputFile{
			TemplateName: "api.go.tmpl",
			OutputPath:   "api/" + strings.ToLower(tag) + ".go",
			Data:         tagOperations{Tag: tag, Operations: tagMap[tag]},
		})
	}

	return files
}
