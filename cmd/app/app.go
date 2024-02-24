package main

import (
	"github.com/khizunov/go-openapi-generator/cmd/bootstrap"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	a := fx.New(
		bootstrap.ProvideConfigs(),
		bootstrap.ProvideLoggers(),
		bootstrap.ProvideServices(),
		bootstrap.ProvideHandlers(),
		bootstrap.ProvideCommands(),

		fx.WithLogger(func(logger *logrus.Logger) fxevent.Logger {
			return &fxevent.ConsoleLogger{
				W: logger.Writer(),
			}
		}),

		fx.Invoke(
			bootstrap.AppendRootCommandHooks,
		),
	)

	a.Run()
}
