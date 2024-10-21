package rally

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
)

type PayRallyFeeUseCase struct {
	base.BaseUseCase
	repo repository.RallyRepository
}

type PayRallyFeeUCParams struct {
	base.UseCaseParams

	Repo repository.RallyRepository
}

func ProvidePayRallyFeeUC(params PayRallyFeeUCParams) *PayRallyFeeUseCase {
	return &PayRallyFeeUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *PayRallyFeeUseCase) Execute(params dto.PayRallyFeeParams) (ipgRedirect string, err error) {
	return "", nil
}
