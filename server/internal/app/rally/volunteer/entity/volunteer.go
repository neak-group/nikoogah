package entity

import "github.com/neak-group/nikoogah/utils/uuid"

type Volunteer struct {
	VolunteerID uuid.UUID
	FullName    string
	Reputation  float32
}


