package app

func RegisterDomainServiceProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if domainServiceProviders == nil {
		domainServiceProviders = append(domainServiceProviders, provider)
	}
}
