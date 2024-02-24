package generate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	enums "github.com/khizunov/go-openapi-generator/internal/definitions/enums"
)

func (c *Command) CMD() *cobra.Command {
	return c.cmd
}

func (c *Command) run(_ *cobra.Command, _ []string) {
	logger := c.log("run")
	logger.Debug("generate command called")

	if e := c.handler.Handle(c.flags.filepath, enums.SpecVersion(c.flags.version)); e != nil {
		logger.Errorf("Failed to generate from %s", c.flags.filepath)
		logger.Error(e)
	} else {
		logger.Info("API is generated successfully")
	}
}

func (c *Command) setFlags() {
	var version string
	var filepath string

	c.cmd.
		Flags().
		StringVarP(&version, "version", "v", "v3", "Spec version of the API")
	c.cmd.
		Flags().
		StringVarP(&filepath, "file", "f", "openapi.yaml", "Path to the OpenAPI/Swagger spec file")

	c.log("setFlags").
		Debugf("Flags are set: version=%s, filepath=%s", version, filepath)

	c.flags = &flags{
		version:  version,
		filepath: filepath,
	}
}

func (c *Command) log(method string) logrus.FieldLogger {
	return c.logger.
		WithField("command", "validate").
		WithField("method", method)
}
