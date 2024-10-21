package entity

import "github.com/neak-group/nikoogah/utils/uuid"

type ParticipationStatus string

const (
	ParticipationPending  ParticipationStatus = "pending"
	ParticipationAccepted ParticipationStatus = "accepted"
	ParticipationRejected ParticipationStatus = "rejected"
)

type HumanParticipation struct {
	VolunteerID uuid.UUID          `bson:"volunteer_id"`
	Phone       string             `bson:"phone"`
	Email       string             `bson:"email"`
	ResumeFile  string             `bson:"resume_file"`
	Status      ParticipationStatus `bson:"status"`
}
