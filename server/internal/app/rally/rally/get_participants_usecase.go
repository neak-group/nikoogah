package rally

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/services"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
)

type GetParticipantsUseCase struct {
	base.BaseUseCase
	repo                        repository.RallyRepository
	rallyParticipationQS        services.RallyParticipationQueryService
	charityRepresentativeAccess services.CharityAccessService
}

type GetParticipantsUCParams struct {
	base.UseCaseParams
	Repo                        repository.RallyRepository
	RallyParticipationQS        services.RallyParticipationQueryService
	CharityRepresentativeAccess services.CharityAccessService
}

func ProvideGetParticipantsUC(params GetParticipantsUCParams) *GetParticipantsUseCase {
	return &GetParticipantsUseCase{
		repo:                        params.Repo,
		rallyParticipationQS:        params.RallyParticipationQS,
		charityRepresentativeAccess: params.CharityRepresentativeAccess,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *GetParticipantsUseCase) Execute(ctx context.Context, params *dto.GetParticipantsParams) ([]*dto.HumanParticipationDTO, error) {
	rally, err := uc.repo.FetchRally(ctx, params.RallyID)
	if err != nil {
		return nil, err
	}

	// TODO[Security]: Check Representative Access
	hasAccess, err := uc.charityRepresentativeAccess.CanViewParticipation(ctx, rally.CharityID)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, fmt.Errorf("use does not have access to this rally participation list")
	}

	participationList, err := uc.rallyParticipationQS.GetRallyHumanParticipation(ctx, rally.ID)
	if err != nil {
		return nil, err
	}

	return participationList, nil
}
