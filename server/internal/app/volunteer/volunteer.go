package volunteer

import "github.com/neak-group/nikoogah/internal/app/volunteer/volunteer"

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	return providers
}

func GetDomainServiceProviders() []interface{} {
	domainServiceProviders := make([]interface{}, 0)

	return domainServiceProviders
}

func GetHandlerProviders() []interface{} {
	handlerProvidors := make([]interface{}, 0)

	handlerProvidors = append(handlerProvidors, volunteer.GetHandlerProviders()...)

	return handlerProvidors
}
