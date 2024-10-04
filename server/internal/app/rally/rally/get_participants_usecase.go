package rally

import (
	"context"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/services"
)

type GetParticipantsUseCase struct {
	app.BaseUseCase
	repo                 repository.RallyRepository
	rallyParticipationQS services.RallyParticipationQueryService
}

type GetParticipantsUCParams struct {
	app.UseCaseParams
	Repo                 repository.RallyRepository
	RallyParticipationQS services.RallyParticipationQueryService
}

func ProvideGetParticipantsUC(params GetParticipantsUCParams) *GetParticipantsUseCase {
	return &GetParticipantsUseCase{
		repo:                 params.Repo,
		rallyParticipationQS: params.RallyParticipationQS,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideGetParticipantsUC)
}

type GetParticipantsParams struct {
	RallyID uuid.UUID
}

func (uc *GetParticipantsUseCase) Execute(ctx context.Context, params GetParticipantsParams) ([]*dto.HumanParticipationDTO, error) {
	rally, err := uc.repo.FetchRally(params.RallyID)
	if err != nil {
		return nil, err
	}

	// TODO[Security]: Check Representative Access

	participationList, err := uc.rallyParticipationQS.GetRallyHumanParticipation(ctx, rally.ID)
	if err != nil {
		return nil, err
	}

	return participationList, nil
}
