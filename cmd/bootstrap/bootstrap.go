package bootstrap

import (
	"go.uber.org/fx"
)

func ProvideCommon() fx.Option {
	return fx.Options(
		ProvideConfigs(),
		ProvideLoggers(),
	)
}
