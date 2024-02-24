package generate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/khizunov/go-openapi-generator/internal/commands/external"
)

type flags struct {
	version  string
	filepath string
}

type Command struct {
	handler external.GenerateHandler
	flags   *flags

	cmd    *cobra.Command
	logger logrus.FieldLogger
}

func NewCommand(handler external.GenerateHandler, logger logrus.FieldLogger) *Command {
	c := &Command{
		cmd: &cobra.Command{
			Use:     "generate",
			Aliases: []string{"gen"},
			Short:   "Generate a client or server from an OpenAPI/Swagger spec",
			Args:    cobra.ExactArgs(1),
		},
		handler: handler,
		logger:  logger,
	}

	c.cmd.Run = c.run

	return c
}
