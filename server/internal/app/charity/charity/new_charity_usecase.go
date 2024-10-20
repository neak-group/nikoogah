package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/events"
	"github.com/neak-group/nikoogah/utils/contextutils"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type RegisterCharityUseCase struct {
	app.BaseUseCase
	repo repository.CharityRepository
}

type RegisterCharityUCParams struct {
	app.UseCaseParams

	Repo repository.CharityRepository
}

func ProvideRegisterCharityUC(params RegisterCharityUCParams) *RegisterCharityUseCase {
	return &RegisterCharityUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRegisterCharityUC)
}

func (uc RegisterCharityUseCase) Execute(ctx context.Context, params *dto.RegisterCharityParams) (uuid.UUID, error) {
	//TODO: fix aggregate transactions

	charity, err := entity.NewCharity(params.Name)
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

	if err := charity.UpdateOfficialData(params.NationalID, params.EconomicID, params.CEO); err != nil {
		return uuid.Nil, err
	}

	repID, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	err = charity.AddRepresentative(repID, entity.Manager)
	if err != nil {
		return uuid.Nil, err
	}

	charityID, err := uc.repo.CreateCharity(ctx, charity)
	if err != nil {
		//TODO: fix error
		return uuid.Nil, err
	}

	charity.Events = append(charity.Events, events.CharityUpdatedEvent{
		ID:            charityID,
		Name:          charity.Name,
		Phone:         charity.Phone.Number,
		Email:         string(charity.EmailAddress),
		MaxRallyLimit: entity.TierMap[charity.CharityTier].GetRallyLimit(),
	})

	//TODO: fire event charity created
	uc.EventDispatcher.DispatchBatch(charity.Events)

	return charityID, nil
}
