package app

func RegisterDomainServiceProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if domainServiceProviders == nil {
		domainServiceProviders = make([]interface{}, 0)
	}
	domainServiceProviders = append(domainServiceProviders, provider)
}
