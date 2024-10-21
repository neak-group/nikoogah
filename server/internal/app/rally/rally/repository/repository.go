package repository

import (
	"context"

	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type RallyRepository interface {
	ApplyFilter(ctx context.Context, filter interface{})
	FetchRally(ctx context.Context, rallyID uuid.UUID) (*entity.Rally, error)
	FetchRallies(ctx context.Context) ([]*entity.Rally, error)
	FetchRalliesByFilter(ctx context.Context, filters ...interface{}) ([]*entity.Rally, error)
	CreateRally(ctx context.Context, rally *entity.Rally) (uuid.UUID, error)
	SaveRally(ctx context.Context, rally *entity.Rally) error
	UpdateParticipations(ctx context.Context, rally *entity.Rally, hp *entity.HumanParticipation, fp *entity.FundParticipation)
	FetchCharityRallyCount(ctx context.Context, charityID uuid.UUID) (int, error)
	FetchRalliesByCharityID(ctx context.Context, charityID uuid.UUID, onlyActive bool) ([]*entity.Rally, error)
	FetchRallyParticipationCount(ctx context.Context, rallyID uuid.UUID, status entity.ParticipationStatus) (int, error)
	FetchTargetFund(ctx context.Context, rallyID uuid.UUID) (decimal.Decimal, error)
}
