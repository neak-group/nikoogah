package app

import (
	"github.com/neak-group/nikoogah/internal/app/rally"
)

func GetDomainServiceProviders() []interface{} {
	domainServiceProviders := make([]interface{}, 0)

	domainServiceProviders = append(domainServiceProviders, rally.GetDomainServiceProviders()...)

	return domainServiceProviders
}
