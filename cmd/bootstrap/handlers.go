package bootstrap

import (
	"go.uber.org/fx"

	"github.com/khizunov/go-openapi-generator/internal/commands/external"
	"github.com/khizunov/go-openapi-generator/internal/handlers/generate"
)

func ProvideHandlers() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				generate.NewHandler,
				fx.As(new(external.GenerateHandler)),
			),
		),
	)
}
