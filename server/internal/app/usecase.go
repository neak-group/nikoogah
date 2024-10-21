package app

import (
	"github.com/neak-group/nikoogah/internal/app/charity"
	"github.com/neak-group/nikoogah/internal/app/rally"
	"github.com/neak-group/nikoogah/internal/app/user"
)

func GetUseCaseProviders() []interface{} {

	useCaseProviders := make([]interface{}, 0)

	useCaseProviders = append(useCaseProviders, user.GetUseCaseProviders()...)
	useCaseProviders = append(useCaseProviders, charity.GetUseCaseProviders()...)
	useCaseProviders = append(useCaseProviders, rally.GetUseCaseProviders()...)

	return useCaseProviders
}
