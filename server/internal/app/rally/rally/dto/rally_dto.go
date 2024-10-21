package dto

import (
	"time"

	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type RallyDTO struct {
	ID                      uuid.UUID       `json:"id"`
	Title                   string          `json:"title"`
	CharityID               uuid.UUID       `json:"charityId"`
	EndDate                 time.Time       `json:"endDate"`
	Description             string          `json:"description"`
	State                   string          `json:"state"`
	RallyFee                decimal.Decimal `json:"rallyFee"`
	NeedsHumanParticipation bool            `json:"needsHumanParticipation"`
	ApplicantCap            int             `json:"applicantCap"`
	HumanParticipationCount int             `json:"humanParticipationCount"`
	NeedsFunding            bool            `json:"needsFunding"`
	FundAmount              decimal.Decimal `json:"fundAmount"`
	OpenFund                bool            `json:"openFund"`
	FundParticipationCount  int             `json:"fundParticipationCount"`
	CreatedAt               time.Time       `json:"createdAt"`
	UpdatedAt               time.Time       `json:"updatedAt"`
}

type FetchRallyParams struct {
	RallyID uuid.UUID
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

type HumanParticipationDTO struct {
	VolunteerID         uuid.UUID
	VolunteerName       string
	VolunteerReputation float32
	Phone               string
	Email               string
	ResumeFile          string
	Status              string
}

type GetParticipantsParams struct {
	RallyID uuid.UUID
}

type FetchCharityRalliesParams struct {
	CharityID  uuid.UUID `json:"charityId"`
	OnlyActive bool      `json:"onlyActive"`
}
