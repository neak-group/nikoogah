package rally

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/dto"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/repository"
)

type NewRallyUseCase struct {
	app.BaseUseCase

	repo repository.RallyRepository
}

type NewRallyUCParams struct {
	app.UseCaseParams

	Repo repository.RallyRepository
}

func ProvideNewRallyUC(params NewRallyUCParams) *NewRallyUseCase {
	return &NewRallyUseCase{
		repo: params.Repo,
		BaseUseCase: app.BaseUseCase{
			Logger:          params.Logger,
			EventDispatcher: params.EventDispatcher,
		},
	}
}

func init() {
	app.RegisterUseCaseProvider(ProvideNewRallyUC)
}

func (uc *NewRallyUseCase) Execute(ctx context.Context, params dto.NewRallyParams) (uuid.UUID, error) {
	max, err := uc.repo.FetchCharityRallyLimit(params.CharityID)
	if err != nil {
		return uuid.Nil, err
	}

	prev, err := uc.repo.FetchRalliesByChrityID(params.CharityID, true)
	if err != nil {
		return uuid.Nil, err
	}

	if len(prev) >= max {
		return uuid.Nil, fmt.Errorf("max rally limit")
	}

	rally, err := entity.NewRally(params.Title, params.Description, params.CharityID, params.EndDate)
	if err != nil {
		return uuid.Nil, err
	}

	// TODO: End Rally Scheduled Job

	err = uc.repo.SaveRally(rally)
	if err != nil {
		return uuid.Nil, err
	}

	uc.EventDispatcher.DispatchBatch(rally.Events)

	return rally.ID, nil
}
