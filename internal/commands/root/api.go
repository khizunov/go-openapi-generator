package root

import (
	"github.com/spf13/cobra"
)

func (c *Command) Execute() error {
	return c.cmd.Execute()
}

func (c *Command) AddCommand(cmd *cobra.Command) {
	c.cmd.AddCommand(cmd)
}
