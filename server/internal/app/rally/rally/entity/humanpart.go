package entity

import "github.com/neak-group/nikoogah/utils/uuid"

type ParticipationStatus string

const (
	ParticipationPending  ParticipationStatus = "pending"
	ParticipationAccepted ParticipationStatus = "accepted"
	ParticipationRejected ParticipationStatus = "rejected"
)

type HumanParticipation struct {
	VolunteerID uuid.UUID
	Phone       string
	Email       string
	ResumeFile  string
	Status      ParticipationStatus
}
