package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type UpdateCredibilityUseCase struct {
	repo CharityRepository
}

type UpdateCredibilityUCParam struct {
	fx.In

	Repo CharityRepository
}

func ProvideUpdateCredibilityUC(params UpdateCredibilityUCParam) *UpdateCredibilityUseCase {
	return &UpdateCredibilityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideUpdateCredibilityUC)
}

type UpdateCredibilityParams struct {
}

func (uc UpdateCredibilityUseCase) Execute(ctx context.Context, params UpdateCredibilityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
