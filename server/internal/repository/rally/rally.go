package rally

import (
	"github.com/neak-group/nikoogah/internal/repository/rally/charityrepo"
	"github.com/neak-group/nikoogah/internal/repository/rally/rallyrepo"
	"github.com/neak-group/nikoogah/internal/repository/rally/volunteerrepo"
)

func GetRepoProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, charityrepo.ProvideMongoRepositoryImpl)
	providers = append(providers, volunteerrepo.ProvideMongoRepositoryImpl)
	providers = append(providers, rallyrepo.ProvideMongoRepositoryImpl)

	return providers
}
