package charitytier

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charitytier/dto"
	"github.com/neak-group/nikoogah/internal/app/charity/charitytier/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type AddCharityTierUseCase struct {
	app.BaseUseCase
	repo CharityTierRepository
}

type AddCharityTierUCParams struct {
	app.UseCaseParams
	Repo CharityTierRepository
}

func ProvideAddCharityTierUC(params AddCharityTierUCParams) *AddCharityTierUseCase {
	return &AddCharityTierUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideAddCharityTierUC)
}

func (uc AddCharityTierUseCase) Execute(ctx context.Context, params dto.AddCharityTierParams) (uuid.UUID, error) {
	ct := entity.NewCharityTier(params.Name, params.RepresentativeLimit)

	err := uc.repo.SaveCharityTier(ct)
	if err != nil {
		return uuid.Nil, err
	}
	return ct.ID, nil
}
