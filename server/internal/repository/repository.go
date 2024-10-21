package repository

import (
	"github.com/neak-group/nikoogah/internal/repository/charity"
	"github.com/neak-group/nikoogah/internal/repository/rally"
	"github.com/neak-group/nikoogah/internal/repository/user"
	"github.com/neak-group/nikoogah/internal/repository/volunteer"
	"go.uber.org/fx"
)

func GetModule() fx.Option {
	return fx.Module("repository", fx.Provide(ProvideRepositories()...))
}

func ProvideRepositories() []interface{} {
	var providers []interface{}

	providers = append(providers, user.GetRepoProviders()...)
	providers = append(providers, charity.GetRepoProviders()...)
	providers = append(providers, rally.GetRepoProviders()...)
	providers = append(providers, volunteer.GetRepoProviders()...)

	return providers
}
