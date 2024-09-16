package charitytier

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charitytier/entity"
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

type AddCharityTierParams struct {
	Name                string
	RepresentativeLimit int
}

func (uc AddCharityTierUseCase) Execute(ctx context.Context, params AddCharityTierParams) (uuid.UUID, error) {
	ct := entity.NewCharityTier(params.Name, params.RepresentativeLimit)

	err := uc.repo.SaveCharityTier(ct)
	if err != nil {
		return uuid.Nil, err
	}
	return ct.ID, nil
}
