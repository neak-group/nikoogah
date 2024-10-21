package entity

import (
	"github.com/neak-group/nikoogah/utils/uuid"
	"github.com/shopspring/decimal"
)

type Volunteer struct {
	UserID                    uuid.UUID       `bson:"user_id"`
	FullName                  string          `bson:"full_name"`
	Reputation                float32         `bson:"reputation"`
	ResumeFile                string          `bson:"resume_file"`
	VolunteerTransactions     int             `bson:"volunteer_transactions"`
	VolunteeringRequestNumber int             `bson:"volunteering_request_number"`
	FinancialAidsSum          decimal.Decimal `bson:"financial_aids_sum"`
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
