package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
)

type ModifyCharityUseCase struct {
	app.BaseUseCase
	repo CharityRepository
}

type ModifyCharityUCParams struct {
	app.UseCaseParams

	Repo CharityRepository
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

type ModifyCharityParams struct {
	ID            uuid.UUID
	Name          string
	Phone         string
	CityPhoneCode string
	Email         string
	Province      string
	City          string
	Address       string
	PostalCode    string
}

func (uc ModifyCharityUseCase) Execute(ctx context.Context, params ModifyCharityParams) (uuid.UUID, error) {
	charity, err := uc.repo.FetchCharity(params.ID)
	if err != nil {
		return uuid.Nil, err
	}
	charity.Events = make([]eventbus.Event, 0)

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

	charityID, err := uc.repo.SaveCharity(charity)
	if err != nil {
		//TODO: fix error
		return uuid.Nil, err
	}

	//TODO: fire event charity created
	uc.EventDispatcher.DispatchBatch(charity.Events)


	return charityID, nil
}
