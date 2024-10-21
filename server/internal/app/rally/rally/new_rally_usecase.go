package rally

import (
	"context"
	"fmt"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/services"
	"github.com/neak-group/nikoogah/internal/core/domain/base"
	"github.com/neak-group/nikoogah/utils/uuid"
)

type NewRallyUseCase struct {
	base.BaseUseCase

	repo                repository.RallyRepository
	charityRallyLimitQS services.CharityRallyQueryService
}

type NewRallyUCParams struct {
	base.UseCaseParams

	Repo                repository.RallyRepository
	CharityRallyLimitQS services.CharityRallyQueryService
}

func ProvideNewRallyUC(params NewRallyUCParams) *NewRallyUseCase {
	return &NewRallyUseCase{
		repo:                params.Repo,
		charityRallyLimitQS: params.CharityRallyLimitQS,
		BaseUseCase: base.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func (uc *NewRallyUseCase) Execute(ctx context.Context, params *dto.NewRallyParams) (uuid.UUID, error) {
	limited, err := uc.charityRallyLimitQS.CheckCharityRallyLimit(ctx, params.CharityID)
	if err != nil {
		return uuid.Nil, err
	}

	if limited {
		return uuid.Nil, fmt.Errorf("rally limit reached")
	}

	rally, err := entity.NewRally(params.Title, params.Description, params.CharityID, params.EndDate)
	if err != nil {
		return uuid.Nil, err
	}

	// TODO: End Rally Scheduled Job

	err = uc.repo.SaveRally(ctx, rally)
	if err != nil {
		return uuid.Nil, err
	}

	uc.EventDispatcher.DispatchBatch(rally.Events)

	return rally.ID, nil
}
