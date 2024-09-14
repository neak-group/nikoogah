package repository

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("repository", fx.Provide(providers...))
}

var providers []interface{}

func RegisterRepositoryProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if providers == nil {
		providers = append(providers, provider)
	}
}
