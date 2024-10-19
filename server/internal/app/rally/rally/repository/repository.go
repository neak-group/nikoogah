package repository

import (
	"github.com/neak-group/nikoogah/internal/app/rally/rally/entity"
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type RallyRepository interface {
	ApplyFilter(filter interface{})
	FetchRally(rallyID uuid.UUID) (*entity.Rally, error)
	FetchRallies() ([]*entity.Rally, error)
	FetchRalliesByFilter(filters ...interface{}) ([]*entity.Rally, error)
	CreateRally(rally *entity.Rally) (uuid.UUID, error)
	SaveRally(rally *entity.Rally) error
	UpdateParticipations(rally *entity.Rally, hp *entity.HumanParticipation, fp *entity.FundParticipation)
	FetchCharityRallyLimit(charityID uuid.UUID) (int, error)
	FetchRalliesByChrityID(charityID uuid.UUID, onlyActive bool) ([]*entity.Rally, error)
	FetchRallyParticipationCount(rallyID uuid.UUID, status entity.ParticipationStatus) (int, error)
	FetchTargetFund(rallyID uuid.UUID) (decimal.Decimal, error)
}
