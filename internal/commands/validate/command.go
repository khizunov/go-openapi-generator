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
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate an OpenAPI/Swagger spec",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("validate command called")
		},
	}
	return &Command{
		cmd:    cmd,
		logger: logger,
	}
}
