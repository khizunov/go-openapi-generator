package v3

import (
	"fmt"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// TransformModels converts all component schemas in a V3 document into codegen.Model values.
func TransformModels(doc *v3high.Document, gen codegen.Generator) ([]codegen.Model, error) {
	if doc.Components == nil || doc.Components.Schemas == nil {
		return nil, nil
	}

	var models []codegen.Model
	for pair := doc.Components.Schemas.First(); pair != nil; pair = pair.Next() {
		name := pair.Key()
		schema := pair.Value()
		if schema == nil {
			continue
		}
		m, err := transformSchema(name, schema, gen)
		if err != nil {
			return nil, fmt.Errorf("schema %q: %w", name, err)
		}
		models = append(models, m)
	}
	return models, nil
}

func transformSchema(name string, schema *base.SchemaProxy, gen codegen.Generator) (codegen.Model, error) {
	s, err := schema.BuildSchema()
	if err != nil {
		return codegen.Model{}, err
	}

	m := codegen.Model{
		Name:         gen.ToModelName(name),
		OriginalName: name,
		Description:  s.Description,
	}

	if len(s.Enum) > 0 {
		m.IsEnum = true
		for _, v := range s.Enum {
			m.EnumValues = append(m.EnumValues, v.Value)
		}
		return m, nil
	}

	schemaType := firstType(s.Type)
	if schemaType == "array" && s.Items != nil && s.Items.IsA() {
		m.IsArray = true
		if itemSchema, err := s.Items.A.BuildSchema(); err == nil {
			m.ArrayItemType = gen.MapType(firstType(itemSchema.Type), itemSchema.Format)
		}
		return m, nil
	}

	if s.Properties == nil {
		return m, nil
	}

	required := toSet(s.Required)
	for pair := s.Properties.First(); pair != nil; pair = pair.Next() {
		propName := pair.Key()
		propProxy := pair.Value()
		if propProxy == nil {
			continue
		}
		propSchema, err := propProxy.BuildSchema()
		if err != nil {
			return codegen.Model{}, fmt.Errorf("property %q: %w", propName, err)
		}

		pt := firstType(propSchema.Type)
		prop := codegen.Property{
			Name:         gen.ToPropertyName(propName),
			OriginalName: propName,
			Type:         gen.MapType(pt, propSchema.Format),
			OpenAPIType:  pt,
			Format:       propSchema.Format,
			Required:     required[propName],
			Description:  propSchema.Description,
		}

		if pt == "array" && propSchema.Items != nil && propSchema.Items.IsA() {
			prop.IsArray = true
			if itemSchema, err := propSchema.Items.A.BuildSchema(); err == nil {
				prop.ArrayItemType = gen.MapType(firstType(itemSchema.Type), itemSchema.Format)
				prop.Type = "[]" + prop.ArrayItemType
			}
		}

		m.Properties = append(m.Properties, prop)
	}
	return m, nil
}

func firstType(types []string) string {
	if len(types) == 0 {
		return ""
	}
	return types[0]
}

func toSet(ss []string) map[string]bool {
	m := make(map[string]bool, len(ss))
	for _, s := range ss {
		m[s] = true
	}
	return m
}
