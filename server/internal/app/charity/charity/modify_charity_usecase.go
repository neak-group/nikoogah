package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type ModifyCharityUseCase struct {
	app.BaseUseCase
	repo repository.CharityRepository
}

type ModifyCharityUCParams struct {
	app.UseCaseParams

	Repo repository.CharityRepository
}

func ProvideModifyCharityUC(params ModifyCharityUCParams) *ModifyCharityUseCase {
	return &ModifyCharityUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideModifyCharityUC)
}

func (uc ModifyCharityUseCase) Execute(ctx context.Context, params dto.ModifyCharityParams) (uuid.UUID, error) {
	charity, err := uc.repo.FetchCharity(ctx, params.ID)
	if err != nil {
		return uuid.Nil, err
	}
	charity.Events = make([]eventbus.Event, 0)

	//TODO[security]: check representative access

	err = charity.UpdateCharityName(params.Name)
	if err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewAddress(params.Province, params.City, params.Address, params.PostalCode); err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewPhone(params.Phone, params.CityPhoneCode); err != nil {
		return uuid.Nil, err
	}

	if err := charity.NewEmail(params.Email); err != nil {
		return uuid.Nil, err
	}

	charityID, err := uc.repo.SaveCharity(ctx, charity)
	if err != nil {
		//TODO: fix error
		return uuid.Nil, err
	}

	//TODO: fire event charity created
	uc.EventDispatcher.DispatchBatch(charity.Events)

	return charityID, nil
}
