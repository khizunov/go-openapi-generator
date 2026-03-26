package generate

import (
	enums "github.com/khizunov/go-openapi-generator/internal/definitions/enums"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (c *Command) CMD() *cobra.Command {
	return c.cmd
}

func (c *Command) run(_ *cobra.Command, _ []string) {
	logger := c.log("run")
	logger.Debug("generate command called")

	if e := c.handler.Handle(c.flags.filepath, enums.SpecVersion(c.flags.version), c.flags.generatorName, c.flags.outputDir); e != nil {
		logger.Errorf("Failed to generate from %s", c.flags.filepath)
		logger.Error(e)
	} else {
		logger.Info("API is generated successfully")
	}
}

func (c *Command) setFlags() {
	c.flags = &flags{}

	c.cmd.Flags().StringVarP(&c.flags.version, "version", "v", "v3", "Spec version of the API")
	c.cmd.Flags().StringVarP(&c.flags.filepath, "file", "f", "openapi.yaml", "Path to the OpenAPI/Swagger spec file")
	c.cmd.Flags().StringVarP(&c.flags.generatorName, "generator", "g", "go-resty", "Generator to use (e.g. go-resty)")
	c.cmd.Flags().StringVarP(&c.flags.outputDir, "output", "o", "./generated", "Output directory")
}

func (c *Command) log(method string) logrus.FieldLogger {
	return c.logger.
		WithField("command", "validate").
		WithField("method", method)
}
