package rally

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
)

type NewHumanParticipationUseCase struct {
	app.BaseUseCase
	repo RallyRepository
}

type NewHumanParticipationUCParams struct {
	app.UseCaseParams

	Repo RallyRepository
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

type NewHumanParticipationParams struct {
	RallyID         uuid.UUID
	VolunteerID     uuid.UUID
	VolunteerPhone  string
	VolunteerEmail  string
	VolunteerResume string
}

func (uc *NewHumanParticipationUseCase) Execute(params NewHumanParticipationParams) error {
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
