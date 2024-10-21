package entity

import (
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type Volunteer struct {
	UserID                    uuid.UUID
	FullName                  string
	Reputation                float32
	ResumeFile                string
	VolunteerTransactions     int
	VolunteeringRequestNumber int
	FinancialAidsSum          decimal.Decimal
}

func UpdateVolunteer(UserID uuid.UUID, FullName string) (*Volunteer, error) {
	return &Volunteer{
		UserID:   UserID,
		FullName: FullName,
	}, nil
}

func (v *Volunteer) UpdateReputation(newReputation float32) error {
	v.Reputation = newReputation

	return nil
}
