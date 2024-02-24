package bootstrap

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"

	"github.com/khizunov/go-openapi-generator/internal/commands/definitions"
	"github.com/khizunov/go-openapi-generator/internal/commands/generate"
	"github.com/khizunov/go-openapi-generator/internal/commands/root"
	"github.com/khizunov/go-openapi-generator/internal/commands/validate"
)

var (
	commandsTag = `group:"commands"`
)

func ProvideCommands() fx.Option {
	return fx.Provide(
		fx.Annotate(
			appendCommands,
			fx.ParamTags(commandsTag),
		),

		AsCommand(generate.NewCommand),
		AsCommand(validate.NewCommand),
	)
}

func appendCommands(commands []definitions.Command, logger logrus.FieldLogger) (*root.Command, error) {
	rc := root.NewCommand(logger)
	for _, c := range commands {
		rc.AddCommand(c.CMD())
	}

	return rc, nil
}

func AppendRootCommandHooks(rc *root.Command, lc fx.Lifecycle, sh fx.Shutdowner) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if e := rc.Execute(); e != nil {
				return sh.Shutdown(fx.ExitCode(1))
			}
			return sh.Shutdown()
		},
	})
}

// AsCommand annotates the given constructor to state that
// it provides a command to the "commands" group.
func AsCommand(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(definitions.Command)),
		fx.ResultTags(commandsTag),
	)
}
