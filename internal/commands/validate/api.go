package validate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (c *Command) CMD() *cobra.Command {
	return c.cmd
}

func (c *Command) run(cmd *cobra.Command, args []string) {
	logger := c.log("run")
	logger.Debug("validate command called")

}

func (c *Command) log(method string) logrus.FieldLogger {
	return c.logger.
		WithField("command", "validate").
		WithField("method", method)
}
