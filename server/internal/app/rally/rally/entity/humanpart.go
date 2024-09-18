package entity

import "github.com/google/uuid"

type ParticipationStatus string

const (
	Pending  ParticipationStatus = "pending"
	Accepted ParticipationStatus = "accepted"
	Rejected ParticipationStatus = "rejected"
)

type HumanParticipation struct {
	VolunteerID uuid.UUID
	Phone       string
	Email       string
	ResumeFile  string
}
