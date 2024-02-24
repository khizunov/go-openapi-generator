package bootstrap

import (
	"go.uber.org/fx"
)

func ProvideConfigs() fx.Option {
	return fx.Provide()
}
