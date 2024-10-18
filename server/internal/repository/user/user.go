package user

import "github.com/neak-group/nikoogah/internal/repository/user/userrepo"

func GetRepoProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, userrepo.ProvideMongoRepositoryImpl)

	return providers
}
