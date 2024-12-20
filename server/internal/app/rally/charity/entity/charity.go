package entity

import "github.com/neak-group/nikoogah/utils/uuid"

type Charity struct {
	CharityID     uuid.UUID
	Name          string
	Phone         string
	Email         string
	MaxRallyLimit int
	IsArchived    bool
}

func UpdateCharity(ID uuid.UUID, name string, phone string, email string) *Charity {
	return &Charity{
		CharityID:  ID,
		Name:       name,
		Phone:      phone,
		Email:      email,
		IsArchived: false,
	}
}
