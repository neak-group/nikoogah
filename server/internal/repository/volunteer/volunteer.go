package volunteer

import "github.com/neak-group/nikoogah/internal/repository/volunteer/volunteerrepo"

func GetRepoProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, volunteerrepo.ProvideMongoRepositoryImpl)

	return providers
}
