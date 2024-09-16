package charity

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"go.uber.org/fx"
)

type AddRepresentativeUseCase struct {
	repo CharityRepository
}

type AddRepresentativeUCParams struct {
	fx.In

	Repo     CharityRepository
	EventBus eventbus.EventBus
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
	CharityID uuid.UUID
	UserID    uuid.UUID
}

func (uc AddRepresentativeUseCase) Execute(ctx context.Context, params AddRepresentativeParams) (uuid.UUID, error) {
	
	return uuid.Nil, nil
}
