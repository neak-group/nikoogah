package app

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application", fx.Provide(providers...))
}

var providers []interface{}

func RegisterUseCaseProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if providers == nil {
		providers = append(providers, provider)
	}
}
