package rally

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"go.uber.org/zap"
)

type FetchRalliesUseCase struct {
	base.BaseUseCase
	repo repository.RallyRepository
}

type FetchRalliesUCParams struct {
	base.UseCaseParams
	Repo repository.RallyRepository
}

func ProvideFetchRalliesUC(params FetchRalliesUCParams) *FetchRalliesUseCase {
	return &FetchRalliesUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *FetchRalliesUseCase) Execute(ctx context.Context) ([]*dto.RallyDTO, error) {
	// Fetch all rallies from the repository
	rallies, err := uc.repo.FetchRallies(ctx)
	if err != nil {
		uc.Logger.Error("Error fetching rallies", zap.Error(err))
		return nil, err
	}

	// Map each rally to a RallyDTO
	var rallyDTOs []*dto.RallyDTO
	for _, rally := range rallies {
		rallyDTO := &dto.RallyDTO{
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
		}
		rallyDTOs = append(rallyDTOs, rallyDTO)
	}

	return rallyDTOs, nil
}
