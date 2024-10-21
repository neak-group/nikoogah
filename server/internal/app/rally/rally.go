package rally

import (
	"github.com/neak-group/nikoogah/internal/app/rally/charity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally"
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer"
)

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, rally.GetUseCaseProviders()...)
	providers = append(providers, volunteer.GetUseCaseProviders()...)

	return providers
}

func GetDomainServiceProviders() []interface{} {
	domainServiceProviders := make([]interface{}, 0)

	return domainServiceProviders
}

func GetHandlerProviders() []interface{} {
	handlerProvidors := make([]interface{}, 0)

	handlerProvidors = append(handlerProvidors, charity.GetHandlerProviders()...)
	handlerProvidors = append(handlerProvidors, volunteer.GetHandlerProviders()...)

	return handlerProvidors
}
