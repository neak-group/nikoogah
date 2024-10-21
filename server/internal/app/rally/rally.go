package rally

import (
	"github.com/neak-group/nikoogah/internal/app/rally/charity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally"
)

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, rally.GetUseCaseProviders()...)

	return providers
}

func GetDomainServiceProviders() []interface{}{
	domainServiceProviders := make([]interface{}, 0)


	return domainServiceProviders
}

func GetHandlerProviders() []interface{} {
	handlerProvidors := make([]interface{}, 0)

	handlerProvidors = append(handlerProvidors, charity.GetHandlerProviders()...)

	return handlerProvidors
}
