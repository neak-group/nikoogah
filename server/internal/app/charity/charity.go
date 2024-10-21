package charity

import "github.com/neak-group/nikoogah/internal/app/charity/charity"


func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, charity.GetUseCaseProviders()...)

	return providers
}

func GetHandlerProviders() []interface{} {
	handlerProviders := make([]interface{}, 0)

	handlerProviders = append(handlerProviders, charity.GetHandlerProviders()...)

	return handlerProviders
}