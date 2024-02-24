package main

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/khizunov/go-openapi-generator/cmd/bootstrap"
)

func main() {
	a := fx.New(
		bootstrap.ProvideCommon(),
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
