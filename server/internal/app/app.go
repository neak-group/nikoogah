package app

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application", fx.Provide(useCaseProviders...), fx.Provide(eventHandlerProviders...))
}

var useCaseProviders []interface{}

var eventHandlerProviders []interface{}

var domainServiceProviders []interface{}
