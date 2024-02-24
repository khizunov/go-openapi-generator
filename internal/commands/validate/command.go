package validate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd    *cobra.Command
	logger logrus.FieldLogger
}

func NewCommand(logger logrus.FieldLogger) *Command {
	c := &Command{
		cmd: &cobra.Command{
			Use:   "validate",
			Short: "Validate an OpenAPI/Swagger spec",
		},
		logger: logger,
	}

	c.setFlags()
	c.cmd.Run = c.run
	return c
}
