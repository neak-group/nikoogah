package entity

import "github.com/google/uuid"

type CharityTier struct {
	ID                  uuid.UUID
	Name                string
	RepresentativeLimit int
}

func NewCharityTier(name string, repLimit int) *CharityTier {
	ct := &CharityTier{
		ID:                  uuid.New(),
		Name:                name,
		RepresentativeLimit: repLimit,
	}

	return ct
}
