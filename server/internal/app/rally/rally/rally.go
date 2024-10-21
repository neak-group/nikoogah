package rally

import "github.com/neak-group/nikoogah/internal/app/rally/rally/services"

func GetUseCaseProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, ProvideNewRallyUC)

	providers = append(providers, ProvidePayRallyFeeUC)

	providers = append(providers, ProvideGetParticipantsUC)

	providers = append(providers, ProvideNewFundParticipationUC)

	providers = append(providers, ProvideNewHumanParticipationUC)

	return providers
}

func GetDomainServiceProviders() []interface{} {
	providers := make([]interface{}, 0)

	providers = append(providers, services.NewRallyParticipationQueryService)

	return providers
}
