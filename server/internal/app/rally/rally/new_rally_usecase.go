package rally

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app"
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/shopspring/decimal"
)

type NewRallyUseCase struct {
	app.BaseUseCase

	repo RallyRepository
}

type NewRallyUCParams struct {
	app.UseCaseParams

	Repo RallyRepository
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

type NewRallyParams struct {
	CharityID   uuid.UUID
	Title       string
	Description string
	EndDate     time.Time

	NeedsFunding bool
	FundAmount   decimal.Decimal

	NeedsHumanParticipation bool
	ApplicantCap            int
}

func (uc *NewRallyUseCase) Execute(ctx context.Context, params NewRallyParams) (uuid.UUID, error) {
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

	for _, e := range rally.Events {
		if err := uc.EventDispatcher.Dispatch(e); err != nil {
			uc.Logger.Error(err.Error())
		}
	}

	return rally.ID, nil
}
