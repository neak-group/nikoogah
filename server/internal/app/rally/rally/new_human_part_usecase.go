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

func ProvideNewHumanPariticipationUC(params NewHumanParticipationUCParams) *NewHumanParticipationUseCase {
	return &NewHumanParticipationUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideNewHumanPariticipationUC)
}

type NewHumanParticipationParams struct {
	RallyID     uuid.UUID
	VolunteerID uuid.UUID
}

func (uc *NewHumanParticipationUseCase) Execute(params NewHumanParticipationParams) error {
	rally, err := uc.repo.FetchRally(params.RallyID)
	if err != nil {
		return err
	}

	count, err:= uc.repo.FetchRallyParticipationCount(rally.ID, entity.Accepted)
	if err != nil {
		return err
	}

	if count>= rally.ApplicantCap {
		return fmt.Errorf("rally limit exceeded")
	}
	

	return nil

}
