package dto

import (
	"github.com/neak-group/nikoogah/utils/uuid"
)

type ProfileDTO struct {
	FullName                  string  `json:"fullName"`
	Reputation                float32 `json:"reputation"`
	ResumeFile                string  `json:"resumeFile"`
	VolunteerTransactions     int     `json:"volunteerTransactions"`
	VolunteeringRequestNumber int     `json:"volunteeringRequestNumber"`
	FinancialAidsSum          string  `json:"financialAidsSum"`
}

type FetchProfileParams struct {
	VolunteerID uuid.UUID `json:"volunteerId"`
}