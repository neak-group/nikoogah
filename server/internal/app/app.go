package app

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application",
		fx.Provide(GetUseCaseProviders()...),
		fx.Provide(GetHandlerProviders()...),
		fx.Provide(GetDomainServiceProviders()...),
	)
}
