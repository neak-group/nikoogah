package charity

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/entity"
	"go.uber.org/fx"
)

type FetchCharityUseCase struct {
	repo CharityRepository
}

type FetchCharityUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideFetchCharityUC(params FetchCharityUCParams) *FetchCharityUseCase {
	return &FetchCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideFetchCharityUC)
}

type FetchCharityParams struct {
}

func (uc FetchCharityUseCase) Execute(ctx context.Context, params FetchCharityParams) (*entity.Charity, error) {
	return nil, nil
}
