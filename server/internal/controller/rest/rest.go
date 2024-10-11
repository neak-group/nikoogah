package rest

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application",
		fx.Option(nil),
	)
}
