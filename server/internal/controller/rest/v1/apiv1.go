package v1

var controllerProviders []any

func RegisterControllerProvider(provider interface{}) {
	if provider == nil {
		return
	}

	if controllerProviders == nil {
		controllerProviders = make([]any, 0)
	}
	controllerProviders = append(controllerProviders, provider)
}

func ProvideV1RestControllers() []any{
	return controllerProviders
}