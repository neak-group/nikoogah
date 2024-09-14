package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
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

func (uc FetchCharityUseCase) Execute(params FetchCharityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
