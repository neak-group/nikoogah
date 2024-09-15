package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/valueobjects"
)

type Charity struct {
	ID             uuid.UUID
	CharityTierID  uuid.UUID
	Name           string
	Address        valueobjects.Address
	Phone          valueobjects.PhoneNumber
	EmailAddress   valueobjects.EmailAddress
	NationalID     string
	EconomicNumber string
	CEO            string

	CreatedAt time.Time
	UpdatedAt time.Time

	Credibility     Credibility
	Representatives []*Representative
}
