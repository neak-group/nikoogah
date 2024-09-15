package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type ModifyCharityUseCase struct {
	repo CharityRepository
}

type ModifyCharityUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideModifyCharityUC(params ModifyCharityUCParams) *ModifyCharityUseCase {
	return &ModifyCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideModifyCharityUC)
}

type ModifyCharityParams struct {
}

func (uc ModifyCharityUseCase) Execute(ctx context.Context, params ModifyCharityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
