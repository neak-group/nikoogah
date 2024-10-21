package charity


func GetHandlerProviders() []interface{} {
	providors := make([]interface{}, 0)

	providors = append(providors, ProvideCharityHandler)

	return providors
}
