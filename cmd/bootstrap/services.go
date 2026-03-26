package bootstrap

import (
	"github.com/khizunov/go-openapi-generator/internal/handlers/generate/external"
	"github.com/khizunov/go-openapi-generator/internal/services/generator"
	"github.com/khizunov/go-openapi-generator/internal/services/parser"
	_ "github.com/khizunov/go-openapi-generator/internal/generators/goresty"
	"go.uber.org/fx"
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
