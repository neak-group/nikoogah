package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
)

type FetchCharityUseCase struct {
	app.BaseUseCase
	repo repository.CharityRepository
}

type FetchCharityUCParams struct {
	app.UseCaseParams

	Repo repository.CharityRepository
}

func ProvideFetchCharityUC(params FetchCharityUCParams) *FetchCharityUseCase {
	return &FetchCharityUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideFetchCharityUC)
}

func (uc FetchCharityUseCase) Execute(ctx context.Context, params *dto.FetchCharityParams) (*dto.CharityDTO, error) {
	charity, err := uc.repo.FetchCharity(ctx, params.CharityID)
	if err != nil {
		return nil, err
	}

	return &dto.CharityDTO{
		Name:       charity.Name,
		Phone:      charity.Phone.Number,
		Email:      string(charity.EmailAddress),
		Address:    charity.Address.String(),
		PostalCode: charity.Address.PostalCode,
		NationalID: charity.NationalID,
		EconomicID: charity.EconomicNumber,
		CEO:        charity.CEO,
	}, nil
}
