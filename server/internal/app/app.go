package app

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application", fx.Provide(providers...), fx.Provide(eventHandlers...))
}

var providers []interface{}

var eventHandlers []interface{}

func RegisterUseCaseProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if providers == nil {
		providers = append(providers, provider)
	}
}

func RegisterHandlerProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if eventHandlers == nil {
		eventHandlers = append(eventHandlers, provider)
	}
}
