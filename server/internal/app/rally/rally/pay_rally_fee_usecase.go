package rally

import (
	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
)

type PayRallyFeeUseCase struct {
	app.BaseUseCase
	repo repository.RallyRepository
}

type PayRallyFeeUCParams struct {
	app.UseCaseParams

	Repo repository.RallyRepository
}

func ProvidePayRallyFeeUC(params PayRallyFeeUCParams) *PayRallyFeeUseCase {
	return &PayRallyFeeUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideNewRallyUC)
}

type PayRallyFeeParams struct {
	RallyID uuid.UUID
}

func (uc *PayRallyFeeUseCase) Execute(params PayRallyFeeParams) (ipgRedirect string, err error) {
	return "", nil
}
