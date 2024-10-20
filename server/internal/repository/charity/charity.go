package charity

import (
	"github.com/neak-group/nikoogah/internal/repository/charity/charityrepo"
)

func GetRepoProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, charityrepo.ProvideMongoRepositoryImpl)

	return providers
}
