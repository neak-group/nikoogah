package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"github.com/shopspring/decimal"
)

type RallyState string

const (
	Pending   RallyState = "pending"
	Active    RallyState = "active"
	Suspended RallyState = "suspended"
	Archived  RallyState = "archived"
)

type Rally struct {
	ID          uuid.UUID
	Title       string
	CharityID   uuid.UUID
	EndDate     time.Time
	Description string
	State       RallyState

	NeedsHumanParticipation bool
	ApplicantCap            int
	HumanParticipations     []*HumanParticipation

	NeedsFunding      bool
	FundAmount        decimal.Decimal
	FundParticipation []*FundParticipation

	CreatedAt time.Time
	UpdatedAt time.Time

	Events []eventbus.Event
}

func NewRally(title, description string, charityID uuid.UUID, EndDate time.Time) (*Rally, error) {
	return &Rally{
		ID:          uuid.New(),
		CharityID:   charityID,
		Title:       title,
		Description: description,
		EndDate:     EndDate,
		State:       Pending,

		HumanParticipations: make([]*HumanParticipation, 0),
		FundParticipation:   make([]*FundParticipation, 0),

		Events: make([]eventbus.Event, 0),
	}, nil
}
