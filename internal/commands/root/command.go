package root

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Command struct {
	cmd    *cobra.Command
	logger logrus.FieldLogger
}

func NewCommand(logger logrus.FieldLogger) *Command {
	cmd := &cobra.Command{
		Use:   "go-openapi-generator",
		Short: "Greate tool to parse, validate OpenAPI/Swagger spec and generate client/server code",
		Run: func(cmd *cobra.Command, args []string) {
			// This function will be executed when the root command is called
			fmt.Println("Welcome to go-openapi-generator! Use --help for usage.")
		},
	}

	return &Command{
		cmd:    cmd,
		logger: logger,
	}
}
