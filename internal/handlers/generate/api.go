package generate

import (
	"fmt"

	enums "github.com/khizunov/go-openapi-generator/internal/definitions/enums"
)

func (h *Handler) Handle(filepath string, specVersion enums.SpecVersion, generatorName, outputDir string) error {
	h.log("Handle").Debugf("Generate from %s using %s → %s", filepath, generatorName, outputDir)

	switch specVersion {
	case enums.SpecVersionV3:
		return h.handleV3(filepath, generatorName, outputDir)
	case enums.SpecVersionV2:
		return h.handleV2(filepath, generatorName, outputDir)
	default:
		return fmt.Errorf("unsupported spec version: %s", specVersion)
	}
}

func (h *Handler) handleV3(filepath, generatorName, outputDir string) error {
	m, e := h.parser.ParseV3(filepath)
	if e != nil {
		return e
	}
	return h.generator.GenerateFromV3(m, generatorName, outputDir)
}

func (h *Handler) handleV2(filepath, generatorName, outputDir string) error {
	m, e := h.parser.ParseV2(filepath)
	if e != nil {
		return e
	}
	return h.generator.GenerateFromV2(m, generatorName, outputDir)
}
