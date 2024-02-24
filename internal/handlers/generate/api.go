package generate

import (
	"fmt"

	enums "github.com/khizunov/go-openapi-generator/internal/definitions/enums"
)

// Handle is a method that handles the generation from an OpenAPI/Swagger spec
func (h *Handler) Handle(filepath string, specVersion enums.SpecVersion) error {
	switch specVersion {
	case enums.SpecVersionV3:
		return h.handleV3(filepath)
	case enums.SpecVersionV2:
		return h.handleV2(filepath)
	default:
		return fmt.Errorf("unsupported spec version: %s", specVersion)
	}
}

func (h *Handler) handleV3(filepath string) error {
	m, e := h.parser.ParseV3(filepath)
	if e != nil {
		return e
	}

	return h.generator.GenerateFromV3(m)
}

func (h *Handler) handleV2(filepath string) error {
	m, e := h.parser.ParseV2(filepath)
	if e != nil {
		return e
	}

	return h.generator.GenerateFromV2(m)
}
