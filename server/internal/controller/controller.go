package controller

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("controller", fx.Provide(providers...))
}

var providers []interface{}

func RegisterControllerProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if providers == nil {
		providers = append(providers, provider)
	}
}
