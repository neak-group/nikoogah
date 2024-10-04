package entity

import "github.com/google/uuid"

type Volunteer struct {
	VolunteerID uuid.UUID
	FullName string
	Reputation float32
}
