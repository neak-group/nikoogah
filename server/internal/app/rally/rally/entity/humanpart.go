package entity

import "github.com/google/uuid"


type HumanParticipation struct{
	VolunteerID uuid.UUID
	Phone string
	Email string
	ResumeFile string	
}