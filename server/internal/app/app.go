package app

import "go.uber.org/fx"

func GetModule() fx.Option {
	return fx.Module("application", fx.Provide(providers...), fx.Provide(eventHandlers...))
}

var providers []interface{}

var eventHandlers []interface{}
