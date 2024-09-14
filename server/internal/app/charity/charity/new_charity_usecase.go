package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type RegisterCharityUseCase struct {
	repo CharityRepository
}

type RegisterCharityUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideRegisterCharityUC(params RegisterCharityUCParams) *RegisterCharityUseCase {
	return &RegisterCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRegisterCharityUC)
}

type RegisterCharityParams struct {
}

func (uc RegisterCharityUseCase) Execute(params RegisterCharityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
