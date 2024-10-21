package volunteer

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/volunteer/repository"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/neak-group/nikoogah/utils/contextutils"
)

type FetchProfileUseCase struct {
	base.BaseUseCase
	repo repository.VolunteerRepository
}

type FetchProfileUCParams struct {
	base.UseCaseParams

	Repo repository.VolunteerRepository
}

func ProvideFetchProfileUC(params FetchProfileUCParams) *FetchProfileUseCase {
	return &FetchProfileUseCase{
		repo: params.Repo,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc FetchProfileUseCase) Execute(ctx context.Context) (*dto.ProfileDTO, error) {
	uid, err := contextutils.GetUserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	volunteer, err := uc.repo.FetchVolunteer(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &dto.ProfileDTO{
		FullName:                  volunteer.FullName,
		Reputation:                volunteer.Reputation,
		ResumeFile:                volunteer.ResumeFile,
		VolunteerTransactions:     volunteer.VolunteerTransactions,
		VolunteeringRequestNumber: volunteer.VolunteerTransactions,
		FinancialAidsSum:          volunteer.FinancialAidsSum.String(),
	}, nil
}
