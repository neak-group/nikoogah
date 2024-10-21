package rally

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
)

type NewHumanParticipationUseCase struct {
	base.BaseUseCase
	repo repository.RallyRepository
}

type NewHumanParticipationUCParams struct {
	base.UseCaseParams

	Repo repository.RallyRepository
}

func ProvideNewHumanParticipationUC(params NewHumanParticipationUCParams) *NewHumanParticipationUseCase {
	return &NewHumanParticipationUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *NewHumanParticipationUseCase) Execute(ctx context.Context, params *dto.NewHumanParticipationParams) error {
	rally, err := uc.repo.FetchRally(ctx, params.RallyID)
	if err != nil {
		return err
	}

	count, err := uc.repo.FetchRallyParticipationCount(ctx, rally.ID, entity.ParticipationAccepted)
	if err != nil {
		return err
	}

	if count >= rally.ApplicantCap {
		return fmt.Errorf("participation limit exceeded")
	}

	err = rally.AddHumanParticipation(params.VolunteerID, params.VolunteerPhone, params.VolunteerEmail, params.VolunteerResume)
	if err != nil {
		return err
	}

	err = uc.repo.SaveRally(ctx, rally)
	if err != nil {
		return err
	}

	uc.EventDispatcher.DispatchBatch(rally.Events)

	return nil

}
