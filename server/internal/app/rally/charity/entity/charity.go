package entity

import "github.com/google/uuid"

type Charity struct {
	CharityID uuid.UUID
	Name      string
	Phone     string
	Email     string
}

func UpdateCharity(ID uuid.UUID, name string, phone string, email string) *Charity{
	return &Charity{
		CharityID:  ID,
		Name: name,
		Phone: phone,
		Email: email,
	}
}