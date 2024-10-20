package rally

import (
	"github.com/neak-group/nikoogah/internal/repository/rally/charityrepo"
)

func GetRepoProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, charityrepo.ProvideMongoRepositoryImpl)

	return providers
}
