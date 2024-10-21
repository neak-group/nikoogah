package rally

import (
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/shopspring/decimal"
)

type NewFundParticipationUseCase struct {
	base.BaseUseCase
	repo repository.RallyRepository
}

type NewFundParticipationUCParams struct {
	base.UseCaseParams

	Repo repository.RallyRepository
}

func ProvideNewFundParticipationUC(params NewFundParticipationUCParams) *NewFundParticipationUseCase {
	return &NewFundParticipationUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *NewFundParticipationUseCase) Execute(params dto.NewFundParticipationParams) error {
	rally, err := uc.repo.FetchRally(params.RallyID)
	if err != nil {
		return err
	}

	var target decimal.Decimal

	if !rally.IsOpenFund() {
		target, err = uc.repo.FetchTargetFund(params.RallyID)
		if err != nil {
			return err
		}
		if params.Amount.Cmp(target) > 0 {
			return fmt.Errorf("funding limit is reached")
		}
	}

	err = rally.AddFundParticipation(params.VolunteerID, params.VolunteerPhone, params.Amount)
	if err != nil {
		return err
	}

	err = uc.repo.SaveRally(rally)
	if err != nil {
		return err
	}

	uc.EventDispatcher.DispatchBatch(rally.Events)

	return nil

}
