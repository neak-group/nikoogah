package charity

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, ProvideRegisterCharityUC)

	providers = append(providers, ProvideModifyCharityUC)

	providers = append(providers, ProvideAddRepresentativeUC)

	providers = append(providers, ProvideCheckRepresentativeAccessUC)

	providers = append(providers, ProvideFetchCharityUC)

	providers = append(providers, ProvideRemoveRepresentativeUC)

	return providers
}

func GetHandlerProviders() []interface{} {
	providors := make([]interface{}, 0)

	return providors
}
