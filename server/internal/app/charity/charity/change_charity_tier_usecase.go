package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type ChangeCharityTierUseCase struct {
	repo CharityRepository
}

type ChangeCharityTierUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideChangeCharityTierUC(params ChangeCharityTierUCParams) *ChangeCharityTierUseCase {
	return &ChangeCharityTierUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideChangeCharityTierUC)
}

type ChangeCharityTierParam struct {
}

func (uc ChangeCharityTierUseCase) Execute(params ChangeCharityTierParam) (uuid.UUID, error) {
	return uuid.Nil, nil
}
