package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
)

type FetchCharityUseCase struct {
	base.BaseUseCase
	repo repository.CharityRepository
}

type FetchCharityUCParams struct {
	base.UseCaseParams

	Repo repository.CharityRepository
}

func ProvideFetchCharityUC(params FetchCharityUCParams) *FetchCharityUseCase {
	return &FetchCharityUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
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
