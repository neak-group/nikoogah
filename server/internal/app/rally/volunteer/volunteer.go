package volunteer

func GetHandlerProviders() []interface{} {
	providors := make([]interface{}, 0)

	providors = append(providors, ProvideVolunteerHandler)

	return providors
}

func GetUseCaseProviders() []interface{} {
	providors := make([]interface{}, 0)

	providors = append(providors, ProvideFetchProfileUC)

	return providors
}
