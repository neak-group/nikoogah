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

func ProvideRegisterCharity(params RegisterCharityUCParams) *RegisterCharityUseCase {
	return &RegisterCharityUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRegisterCharity)
}

type NewCharityParams struct {
}

func (uc RegisterCharityUseCase) Execute(params NewCharityParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
