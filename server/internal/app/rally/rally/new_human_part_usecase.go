package rally

import (
	"fmt"

	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
)

type NewHumanParticipationUseCase struct {
	app.BaseUseCase
	repo repository.RallyRepository
}

type NewHumanParticipationUCParams struct {
	app.UseCaseParams

	Repo repository.RallyRepository
}

func ProvideNewHumanParticipationUC(params NewHumanParticipationUCParams) *NewHumanParticipationUseCase {
	return &NewHumanParticipationUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideNewHumanParticipationUC)
}

func (uc *NewHumanParticipationUseCase) Execute(params dto.NewHumanParticipationParams) error {
	rally, err := uc.repo.FetchRally(params.RallyID)
	if err != nil {
		return err
	}

	count, err := uc.repo.FetchRallyParticipationCount(rally.ID, entity.ParticipationAccepted)
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

	err = uc.repo.SaveRally(rally)
	if err != nil {
		return err
	}

	uc.EventDispatcher.DispatchBatch(rally.Events)

	return nil

}
