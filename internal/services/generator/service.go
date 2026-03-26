package generator

import (
	"fmt"
	"path/filepath"

	"github.com/khizunov/go-openapi-generator/internal/codegen"
	"github.com/khizunov/go-openapi-generator/internal/renderer"
	transformerv3 "github.com/khizunov/go-openapi-generator/internal/transformer/v3"
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
	"github.com/sirupsen/logrus"
)

type Generator struct {
	logger logrus.FieldLogger
}

func NewGenerator(logger logrus.FieldLogger) *Generator {
	return &Generator{logger: logger}
}

func (g *Generator) GenerateFromV3(doc *v3.Document, generatorName, outputDir string) error {
	log := g.logger.WithField("generator", generatorName).WithField("output", outputDir)

	gen, err := codegen.Lookup(generatorName)
	if err != nil {
		return err
	}

	log.Info("transforming models")
	models, err := transformerv3.TransformModels(doc, gen)
	if err != nil {
		return err
	}

	log.Info("transforming operations")
	ops, err := transformerv3.TransformOperations(doc, gen)
	if err != nil {
		return err
	}

	log.Infof("rendering %d models, %d operations", len(models), len(ops))
	for _, f := range gen.OutputFiles(models, ops) {
		outPath := filepath.Join(outputDir, f.OutputPath)
		log.Debugf("rendering %s → %s", f.TemplateName, outPath)
		if err := renderer.Render(gen.TemplateFS(), f.TemplateName, f.Data, outPath); err != nil {
			return err
		}
	}
	return nil
}

func (g *Generator) GenerateFromV2(_ *v2.Swagger, _, _ string) error {
	return fmt.Errorf("V2 generation not yet implemented")
}
