package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type AddRepresentativeUseCase struct {
	repo CharityRepository
}

type AddRepresentativeUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideAddRepresentativeUC(params AddRepresentativeUCParams) *AddRepresentativeUseCase {
	return &AddRepresentativeUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideAddRepresentativeUC)
}

type AddRepresentativeParams struct {
}

func (uc AddRepresentativeUseCase) Execute(ctx context.Context, params AddRepresentativeParams) (uuid.UUID, error) {
	
	return uuid.Nil, nil
}
