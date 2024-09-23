package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/core/service/eventbus"
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
	ID          uuid.UUID
	Title       string
	CharityID   uuid.UUID
	EndDate     time.Time
	Description string
	State       RallyState
	RallyFee    decimal.Decimal

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
		State:       PendingVerification,

		HumanParticipations: make([]*HumanParticipation, 0),
		FundParticipation:   make([]*FundParticipation, 0),

		Events: make([]eventbus.Event, 0),
	}, nil
}


func (r *Rally) AddHumanParticipation( volunteerID uuid.UUID, volunteerPhone string, volunteerEmail string, resumePath string) error{
	//Validate Phone

	//Validate Email

	r.HumanParticipations = append(r.HumanParticipations, &HumanParticipation{
		VolunteerID: volunteerID,
		Phone: volunteerPhone,
		Email: volunteerEmail,
		ResumeFile: resumePath,
		Status: ParticipationAccepted,
	})
	
	return nil
}