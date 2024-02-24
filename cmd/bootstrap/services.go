package bootstrap

import (
	"go.uber.org/fx"

	"github.com/khizunov/go-openapi-generator/internal/handlers/generate/external"
	"github.com/khizunov/go-openapi-generator/internal/services/generator"
	"github.com/khizunov/go-openapi-generator/internal/services/parser"
)

func ProvideServices() fx.Option {
	return fx.Provide(
		fx.Annotate(
			parser.NewParser,
			fx.As(new(external.ParserAPI)),
		),
		fx.Annotate(
			generator.NewGenerator,
			fx.As(new(external.GeneratorAPI)),
		),
	)
}
