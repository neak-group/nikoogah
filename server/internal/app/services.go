package app

import (
	"github.com/neak-group/nikoogah/internal/app/rally"
	"github.com/neak-group/nikoogah/internal/app/volunteer"
)

func GetDomainServiceProviders() []interface{} {
	domainServiceProviders := make([]interface{}, 0)

	domainServiceProviders = append(domainServiceProviders, rally.GetDomainServiceProviders()...)
	domainServiceProviders = append(domainServiceProviders, volunteer.GetDomainServiceProviders()...)

	return domainServiceProviders
}
