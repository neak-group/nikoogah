package repository

import (
	"github.com/neak-group/nikoogah/internal/repository/charity"
	"github.com/neak-group/nikoogah/internal/repository/user"
	"go.uber.org/fx"
)

func GetModule() fx.Option {
	return fx.Module("repository", fx.Provide(ProvideRepositories()...))
}

func ProvideRepositories() []interface{} {
	var providers []interface{}

	providers = append(providers, user.GetRepoProviders()...)
	providers = append(providers, charity.GetRepoProviders()...)

	return providers
}
