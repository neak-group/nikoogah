package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
)

type FetchCharityUseCase struct {
	app.BaseUseCase
	repo CharityRepository
}

type FetchCharityUCParams struct {
	app.UseCaseParams

	Repo CharityRepository
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

type FetchCharityParams struct {
	CharityID uuid.UUID
}

func (uc FetchCharityUseCase) Execute(ctx context.Context, params FetchCharityParams) (*entity.Charity, error) {
	charity, err := uc.repo.FetchCharity(params.CharityID)
	if err != nil {
		return nil, err
	}

	return charity, nil
}
