package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/neak-group/nikoogah/internal/app/charity/charity/valueobjects"
)

type Charity struct {
	CharityTierID  uuid.UUID
	Name           string
	Address        string
	Phone          valueobjects.PhoneNumber
	EmailAddress   valueobjects.EmailAddress
	PostalCode     string
	NationalID     string
	EconomicNumber string
	CEO            string

	CreatedAt time.Time
	UpdatedAt time.Time

	Credibility     Credibility
	Representatives []*Representative
}
