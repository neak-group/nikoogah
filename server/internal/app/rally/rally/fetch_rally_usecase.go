package rally

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"go.uber.org/zap"
)

type FetchRallyUseCase struct {
	base.BaseUseCase
	repo repository.RallyRepository
}

type FetchRallyUCParams struct {
	base.UseCaseParams
	Repo repository.RallyRepository
}

func ProvideFetchRallyUC(params FetchRallyUCParams) *FetchRallyUseCase {
	return &FetchRallyUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *FetchRallyUseCase) Execute(ctx context.Context, params *dto.FetchRallyParams) (*dto.RallyDTO, error) {
	// Fetch the rally from the repository using the provided rally ID
	rally, err := uc.repo.FetchRally(ctx, params.RallyID)
	if err != nil {
		uc.Logger.Error("Error fetching rally", zap.Error(err))
		return nil, err
	}

	// If the rally is not found, return an error
	if rally == nil {
		return nil, fmt.Errorf("rally with ID %s not found", params.RallyID)
	}

	// Map the rally entity to the RallyDTO for response
	return &dto.RallyDTO{
		ID:                      rally.ID,
		Title:                   rally.Title,
		CharityID:               rally.CharityID,
		EndDate:                 rally.EndDate,
		Description:             rally.Description,
		State:                   string(rally.State),
		RallyFee:                rally.RallyFee,
		NeedsHumanParticipation: rally.NeedsHumanParticipation,
		ApplicantCap:            rally.ApplicantCap,
		HumanParticipationCount: len(rally.HumanParticipations),
		NeedsFunding:            rally.NeedsFunding,
		FundAmount:              rally.FundAmount,
		OpenFund:                rally.OpenFund,
		FundParticipationCount:  len(rally.FundParticipation),
		CreatedAt:               rally.CreatedAt,
		UpdatedAt:               rally.UpdatedAt,
	}, nil
}
