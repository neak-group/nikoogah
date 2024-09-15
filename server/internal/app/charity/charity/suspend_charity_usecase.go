package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type SuspendCharityUseCase struct {
	repo CharityRepository
}

type SuspendCharityUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideSuspendCharityUC(params SuspendCharityUCParams) *SuspendCharityUseCase {
	return &SuspendCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideSuspendCharityUC)
}

type SuspendCharityParams struct {
}

func (uc SuspendCharityUseCase) Execute(ctx context.Context, params SuspendCharityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
