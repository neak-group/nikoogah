package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

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


type NewHumanParticipationParams struct {
	RallyID         uuid.UUID
	VolunteerID     uuid.UUID
	VolunteerPhone  string
	VolunteerEmail  string
	VolunteerResume string
}

type NewFundParticipationParams struct {
	RallyID        uuid.UUID
	VolunteerID    uuid.UUID
	VolunteerPhone string
	Amount         decimal.Decimal
}

type PayRallyFeeParams struct {
	RallyID uuid.UUID
}