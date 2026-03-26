package v3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	v3high "github.com/pb33f/libopenapi/datamodel/high/v3"
)

// TransformOperations converts all paths in a V3 document into codegen.Operation values.
func TransformOperations(doc *v3high.Document, gen codegen.Generator) ([]codegen.Operation, error) {
	if doc.Paths == nil || doc.Paths.PathItems == nil {
		return nil, nil
	}

	var ops []codegen.Operation
	for pair := doc.Paths.PathItems.First(); pair != nil; pair = pair.Next() {
		path := pair.Key()
		item := pair.Value()
		if item == nil {
			continue
		}
		itemOps, err := pathItemOperations(path, item, gen)
		if err != nil {
			return nil, fmt.Errorf("path %q: %w", path, err)
		}
		ops = append(ops, itemOps...)
	}
	return ops, nil
}

func pathItemOperations(path string, item *v3high.PathItem, gen codegen.Generator) ([]codegen.Operation, error) {
	type methodOp struct {
		method string
		op     *v3high.Operation
	}
	candidates := []methodOp{
		{"GET", item.Get},
		{"POST", item.Post},
		{"PUT", item.Put},
		{"DELETE", item.Delete},
		{"PATCH", item.Patch},
		{"OPTIONS", item.Options},
		{"HEAD", item.Head},
	}

	var ops []codegen.Operation
	for _, c := range candidates {
		if c.op == nil {
			continue
		}
		op, err := transformOperation(path, c.method, c.op, gen)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}
	return ops, nil
}

func transformOperation(path, method string, op *v3high.Operation, gen codegen.Generator) (codegen.Operation, error) {
	result := codegen.Operation{
		OriginalID:  op.OperationId,
		Name:        gen.ToOperationName(op.OperationId),
		Method:      method,
		Path:        path,
		Summary:     op.Summary,
		Description: op.Description,
		Tags:        op.Tags,
	}

	for _, paramProxy := range op.Parameters {
		if paramProxy == nil {
			continue
		}
		p := paramProxy
		pt := ""
		format := ""
		if p.Schema != nil {
			if s, err := p.Schema.BuildSchema(); err == nil {
				pt = firstType(s.Type)
				format = s.Format
			}
		}
		required := p.Required != nil && *p.Required
		result.Parameters = append(result.Parameters, codegen.Parameter{
			Name:         gen.ToParamName(p.Name),
			OriginalName: p.Name,
			In:           codegen.ParameterIn(p.In),
			Type:         gen.MapType(pt, format),
			OpenAPIType:  pt,
			Format:       format,
			Required:     required,
			Description:  p.Description,
		})
	}

	result.PathFmt, result.PathParamNames = buildPathFmt(result.Path, result.Parameters)

	if op.RequestBody != nil {
		bodyType := resolveRequestBodyType(op.RequestBody, gen)
		bodyRequired := op.RequestBody.Required != nil && *op.RequestBody.Required
		result.RequestBody = &codegen.Parameter{
			Name:     "body",
			In:       codegen.ParameterInPath,
			Type:     bodyType,
			IsBody:   true,
			Required: bodyRequired,
		}
	}

	result.SuccessResponse = resolveSuccessResponse(op, gen)
	if result.SuccessResponse != nil {
		result.Responses = []codegen.Response{*result.SuccessResponse}
	}

	if result.Name == "" {
		result.Name = gen.ToOperationName(method + "_" + sanitizePath(path))
		result.OriginalID = result.Name
	}

	return result, nil
}

func resolveRequestBodyType(rb *v3high.RequestBody, gen codegen.Generator) string {
	if rb.Content == nil {
		return ""
	}
	mt := rb.Content.GetOrZero("application/json")
	if mt == nil || mt.Schema == nil {
		return ""
	}
	s, err := mt.Schema.BuildSchema()
	if err != nil {
		return ""
	}
	return gen.MapType(firstType(s.Type), s.Format)
}

func resolveSuccessResponse(op *v3high.Operation, gen codegen.Generator) *codegen.Response {
	if op.Responses == nil || op.Responses.Codes == nil {
		return nil
	}
	for _, code := range []string{"200", "201", "202"} {
		resp := op.Responses.Codes.GetOrZero(code)
		if resp == nil {
			continue
		}
		r := &codegen.Response{
			StatusCode:  code,
			Description: resp.Description,
		}
		if resp.Content != nil {
			mt := resp.Content.GetOrZero("application/json")
			if mt != nil && mt.Schema != nil {
				if s, err := mt.Schema.BuildSchema(); err == nil {
					t := firstType(s.Type)
					r.IsArray = t == "array"
					r.Type = gen.MapType(t, s.Format)
				}
			}
		}
		return r
	}
	return nil
}

func buildPathFmt(path string, params []codegen.Parameter) (string, []string) {
	paramTypes := map[string]string{}
	for _, p := range params {
		if p.In == codegen.ParameterInPath {
			paramTypes[p.OriginalName] = p.OpenAPIType
		}
	}

	var names []string
	result := regexp.MustCompile(`\{(\w+)\}`).ReplaceAllStringFunc(path, func(m string) string {
		name := m[1 : len(m)-1]
		names = append(names, name)
		if t := paramTypes[name]; t == "integer" {
			return "%d"
		}
		return "%s"
	})
	return result, names
}

func sanitizePath(path string) string {
	r := strings.NewReplacer("/", "_", "{", "", "}", "")
	return strings.Trim(r.Replace(path), "_")
}
