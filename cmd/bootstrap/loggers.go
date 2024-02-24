package bootstrap

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func ProvideLoggers() fx.Option {
	return fx.Provide(
		logrus.New,
		fx.Annotate(
			logrus.New,
			fx.As(new(logrus.FieldLogger)),
		),
	)
}
