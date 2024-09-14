package charity

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"go.uber.org/fx"
)

type RemoveRepresentativeUseCase struct {
	repo CharityRepository
}

type RemoveRepresentativeUCParams struct {
	fx.In

	Repo CharityRepository
}

func ProvideRemoveRepresentativeUC(params RemoveRepresentativeUCParams) *RemoveRepresentativeUseCase {
	return &RemoveRepresentativeUseCase{
		repo: params.Repo,
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideRemoveRepresentativeUC)
}

type RemoveRepresentativeParams struct {
}

func (uc RemoveRepresentativeUseCase) Execute(params RemoveRepresentativeParams) (uuid.UUID, error) {
	return uuid.Nil, nil
}
