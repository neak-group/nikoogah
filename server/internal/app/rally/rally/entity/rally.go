package entity

import (
	"fmt"
	"time"

	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type RallyState string

const (
	PendingVerification RallyState = "pending-verification"
	Active              RallyState = "active"
	PendingReport       RallyState = "pending-report"
	Suspended           RallyState = "suspended"
	Archived            RallyState = "archived"
)

type Rally struct {
	ID          uuid.UUID          `bson:"id"`
	Title       string             `bson:"title"`
	CharityID   uuid.UUID          `bson:"charity_id"`
	EndDate     time.Time          `bson:"end_date"`
	Description string             `bson:"description"`
	State       RallyState         `bson:"state"`
	RallyFee    decimal.Decimal    `bson:"rally_fee"`

	NeedsHumanParticipation bool              `bson:"needs_human_participation"`
	ApplicantCap            int               `bson:"applicant_cap"`
	HumanParticipations     []*HumanParticipation `bson:"human_participations"`

	NeedsFunding      bool              `bson:"needs_funding"`
	FundAmount        decimal.Decimal   `bson:"fund_amount"`
	OpenFund          bool              `bson:"open_fund"`
	FundParticipation []*FundParticipation `bson:"fund_participation"`

	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`

	Events []eventbus.Event      `bson:"-"`
}

func NewRally(title, description string, charityID uuid.UUID, EndDate time.Time) (*Rally, error) {
	return &Rally{
		ID:          uuid.New(),
		CharityID:   charityID,
		Title:       title,
		Description: description,
		EndDate:     EndDate,
		State:       PendingVerification,

		HumanParticipations: make([]*HumanParticipation, 0),
		FundParticipation:   make([]*FundParticipation, 0),

		Events: make([]eventbus.Event, 0),
	}, nil
}

func (r *Rally) IsOpenFund() bool {
	return r.OpenFund
}

func (r *Rally) AddHumanParticipation(volunteerID uuid.UUID, volunteerPhone string, volunteerEmail string, resumePath string) error {
	//Validate Phone

	//Validate Email

	if r.ApplicantCap == 0 {
		return fmt.Errorf("no need for human participation")
	}

	r.HumanParticipations = append(r.HumanParticipations, &HumanParticipation{
		VolunteerID: volunteerID,
		Phone:       volunteerPhone,
		Email:       volunteerEmail,
		ResumeFile:  resumePath,
		Status:      ParticipationAccepted,
	})

	return nil
}

func (r *Rally) AddFundParticipation(volunteerID uuid.UUID, volunteerPhone string, amount decimal.Decimal) error {
	if r.FundAmount.IsZero() {
		return fmt.Errorf("no need for fund participation")
	}

	r.FundParticipation = append(r.FundParticipation, &FundParticipation{
		VolunteerID: volunteerID,
		Amount:      amount,
		Date:        time.Now(),
	})

	return nil
}
